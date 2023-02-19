package subfunction

import (
	"context"
	api_input_reader "data-platform-api-delivery-document-items-creates-subfunc/API_Input_Reader"
	dpfm_api_output_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Output_Formatter"
	api_processing_data_formatter "data-platform-api-delivery-document-items-creates-subfunc/API_Processing_Data_Formatter"
	"data-platform-api-delivery-document-items-creates-subfunc/config"
	"strings"

	"sync"

	"github.com/latonaio/golang-logging-library-for-data-platform/logger"
	database "github.com/latonaio/golang-mysql-network-connector"
	rabbitmq "github.com/latonaio/rabbitmq-golang-client-for-data-platform"
	"golang.org/x/xerrors"
)

type SubFunction struct {
	ctx  context.Context
	db   *database.Mysql
	conf *config.Conf
	rmq  *rabbitmq.RabbitmqClient
	l    *logger.Logger
}

func NewSubFunction(ctx context.Context, db *database.Mysql, conf *config.Conf, rmq *rabbitmq.RabbitmqClient, l *logger.Logger) *SubFunction {
	return &SubFunction{
		ctx:  ctx,
		db:   db,
		conf: conf,
		rmq:  rmq,
		l:    l,
	}
}

func (f *SubFunction) MetaData(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.MetaData {
	metaData := psdc.ConvertToMetaData(sdc)

	return metaData
}

func (f *SubFunction) ProcessType(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) *api_processing_data_formatter.ProcessType {
	processType := psdc.ConvertToProcessType()

	processType.BulkProcess = true
	// processType.IndividualProcess = true

	if processType.BulkProcess {
		// processType.ArraySpec = true
		processType.RangeSpec = true
	}

	return processType
}

func (f *SubFunction) OrderItemInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderItem, error) {
	data := make([]*api_processing_data_formatter.OrderItem, 0)
	var err error

	processType := psdc.ProcessType

	if processType.ArraySpec {
		data, err = f.OrderItemByArraySpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else if processType.RangeSpec {
		data, err = f.OrderItemByRangeSpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (f *SubFunction) OrderItemByArraySpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderItem, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToOrderItemKey()

	deliverToParty := sdc.InputParameters.DeliverToParty
	deliverFromParty := sdc.InputParameters.DeliverFromParty
	deliverToPlant := sdc.InputParameters.DeliverToPlant
	deliverFromPlant := sdc.InputParameters.DeliverFromPlant

	for i := range *deliverToParty {
		dataKey.DeliverToParty = append(dataKey.DeliverToParty, (*deliverToParty)[i])
	}
	for i := range *deliverFromParty {
		dataKey.DeliverFromParty = append(dataKey.DeliverFromParty, (*deliverFromParty)[i])
	}
	for i := range *deliverToPlant {
		dataKey.DeliverToPlant = append(dataKey.DeliverToPlant, (*deliverToPlant)[i])
	}
	for i := range *deliverFromPlant {
		dataKey.DeliverFromPlant = append(dataKey.DeliverFromPlant, (*deliverFromPlant)[i])
	}

	repeat1 := strings.Repeat("?,", len(dataKey.DeliverToParty)-1) + "?"
	for _, v := range dataKey.DeliverToParty {
		args = append(args, v)
	}
	repeat2 := strings.Repeat("?,", len(dataKey.DeliverFromParty)-1) + "?"
	for _, v := range dataKey.DeliverFromParty {
		args = append(args, v)
	}
	repeat3 := strings.Repeat("?,", len(dataKey.DeliverToPlant)-1) + "?"
	for _, v := range dataKey.DeliverToPlant {
		args = append(args, v)
	}
	repeat4 := strings.Repeat("?,", len(dataKey.DeliverFromPlant)-1) + "?"
	for _, v := range dataKey.DeliverFromPlant {
		args = append(args, v)
	}

	args = append(args, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemBlockStatus, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus)

	count := new(int)
	err := f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE DeliverToParty IN ( `+repeat1+` )
		AND DeliverFromParty IN ( `+repeat2+` )
		AND DeliverToPlant IN ( `+repeat3+` )
		AND DeliverFromPlant IN ( `+repeat4+` )
		AND (ItemCompleteDeliveryIsDefined, ItemBlockStatus, ItemDeliveryBlockStatus) = (?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, args...,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderID, OrderItemの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant,
		ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemBlockStatus, ItemDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE DeliverToParty IN ( `+repeat1+` )
		AND DeliverFromParty IN ( `+repeat2+` )
		AND DeliverToPlant IN ( `+repeat3+` )
		AND DeliverFromPlant IN ( `+repeat4+` )
		AND (ItemCompleteDeliveryIsDefined, ItemBlockStatus, ItemDeliveryBlockStatus) = (?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderItem(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderItemByRangeSpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderItem, error) {
	dataKey := psdc.ConvertToOrderItemKey()

	dataKey.DeliverToPartyFrom = sdc.InputParameters.DeliverToPartyFrom
	dataKey.DeliverToPartyTo = sdc.InputParameters.DeliverToPartyTo
	dataKey.DeliverFromPartyFrom = sdc.InputParameters.DeliverFromPartyFrom
	dataKey.DeliverFromPartyTo = sdc.InputParameters.DeliverFromPartyTo
	dataKey.DeliverToPlantFrom = sdc.InputParameters.DeliverToPlantFrom
	dataKey.DeliverToPlantTo = sdc.InputParameters.DeliverToPlantTo
	dataKey.DeliverFromPlantFrom = sdc.InputParameters.DeliverFromPlantFrom
	dataKey.DeliverFromPlantTo = sdc.InputParameters.DeliverFromPlantTo

	count := new(int)
	err := f.db.QueryRow(
		`SELECT COUNT(*)
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE DeliverToParty BETWEEN ? AND ?
		AND DeliverFromParty BETWEEN ? AND ?
		AND DeliverToPlant BETWEEN ? AND ?
		AND DeliverFromPlant BETWEEN ? AND ?
		AND (ItemCompleteDeliveryIsDefined, ItemBlockStatus, ItemDeliveryBlockStatus) = (?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, dataKey.DeliverToPartyFrom, dataKey.DeliverToPartyTo, dataKey.DeliverFromPartyFrom, dataKey.DeliverFromPartyTo, dataKey.DeliverToPlantFrom, dataKey.DeliverToPlantTo, dataKey.DeliverFromPlantFrom, dataKey.DeliverFromPlantTo, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemBlockStatus, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus,
	).Scan(&count)
	if err != nil {
		return nil, err
	}
	if *count == 0 || *count > 1000 {
		return nil, xerrors.Errorf("OrderID, OrderItemの検索結果がゼロ件または1,000件超です。")
	}

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, DeliverToParty, DeliverFromParty, DeliverToPlant, DeliverFromPlant,
		ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemBlockStatus, ItemDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE DeliverToParty BETWEEN ? AND ?
		AND DeliverFromParty BETWEEN ? AND ?
		AND DeliverToPlant BETWEEN ? AND ?
		AND DeliverFromPlant BETWEEN ? AND ?
		AND (ItemCompleteDeliveryIsDefined, ItemBlockStatus, ItemDeliveryBlockStatus) = (?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, dataKey.DeliverToPartyFrom, dataKey.DeliverToPartyTo, dataKey.DeliverFromPartyFrom, dataKey.DeliverFromPartyTo, dataKey.DeliverToPlantFrom, dataKey.DeliverToPlantTo, dataKey.DeliverFromPlantFrom, dataKey.DeliverFromPlantTo, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemBlockStatus, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderItem(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrderItemInIndividualProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrderItem, error) {
	dataKey := psdc.ConvertToOrderItemInIndividualProcessKey()

	dataKey.OrderID = sdc.Header.ReferenceDocument
	dataKey.OrderItem = sdc.Header.ReferenceDocumentItem

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, ItemCompleteDeliveryIsDefined, ItemDeliveryStatus, ItemDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_data
		WHERE (OrderID, OrderItem, ItemCompleteDeliveryIsDefined, ItemDeliveryBlockStatus) = (?, ?, ?, ?)
		AND ItemDeliveryStatus <> ?;`, dataKey.OrderID, dataKey.OrderItem, dataKey.ItemCompleteDeliveryIsDefined, dataKey.ItemDeliveryBlockStatus, dataKey.ItemDeliveryStatus,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrderItemInIndividualProcess(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrdersItemScheduleLineInBulkProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersItemScheduleLine, error) {
	data := make([]*api_processing_data_formatter.OrdersItemScheduleLine, 0)
	var err error

	processType := psdc.ProcessType

	if processType.ArraySpec {
		data, err = f.OrdersItemScheduleLineByArraySpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	} else if processType.RangeSpec {
		data, err = f.OrdersItemScheduleLineByRangeSpec(sdc, psdc)
		if err != nil {
			return nil, err
		}
	}

	return data, nil
}

func (f *SubFunction) OrdersItemScheduleLineByArraySpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersItemScheduleLine, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToOrdersItemScheduleLineKey()

	orderItem := psdc.OrderItem
	confirmedDeliveryDate := sdc.InputParameters.ConfirmedDeliveryDate

	for i := range orderItem {
		dataKey.OrderID = append(dataKey.OrderID, (orderItem)[i].OrderID)
		dataKey.OrderItem = append(dataKey.OrderItem, (orderItem)[i].OrderItem)
	}
	for i := range *confirmedDeliveryDate {
		dataKey.ConfirmedDeliveryDate = append(dataKey.ConfirmedDeliveryDate, (*confirmedDeliveryDate)[i])
	}

	repeat1 := strings.Repeat("(?,?),", len(dataKey.OrderID)-1) + "(?,?)"
	for i := range dataKey.OrderID {
		args = append(args, dataKey.OrderID[i], dataKey.OrderItem[i])
	}
	repeat2 := strings.Repeat("?,", len(dataKey.ConfirmedDeliveryDate)-1) + "?"
	for _, v := range dataKey.ConfirmedDeliveryDate {
		args = append(args, v)
	}

	args = append(args, dataKey.ItemScheduleLineDeliveryBlockStatus, dataKey.OpenConfirmedQuantityInBaseUnit)

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, ScheduleLine, RequestedDeliveryDate, ConfirmedDeliveryDate, OrderQuantityInBaseUnit,
		ConfirmedOrderQuantityByPDTAvailCheck, OpenConfirmedQuantityInBaseUnit, StockIsFullyConfirmed, ItemScheduleLineDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
		WHERE (OrderID, OrderItem) IN ( `+repeat1+` )
		AND ConfirmedDeliveryDate IN ( `+repeat2+` )
		AND ItemScheduleLineDeliveryBlockStatus = ?
		AND OpenConfirmedQuantityInBaseUnit <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersItemScheduleLine(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrdersItemScheduleLineByRangeSpec(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersItemScheduleLine, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToOrdersItemScheduleLineKey()

	orderItem := psdc.OrderItem

	for i := range orderItem {
		dataKey.OrderID = append(dataKey.OrderID, (orderItem)[i].OrderID)
		dataKey.OrderItem = append(dataKey.OrderItem, (orderItem)[i].OrderItem)
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.OrderID)-1) + "(?,?)"
	for i := range dataKey.OrderID {
		args = append(args, dataKey.OrderID[i], dataKey.OrderItem[i])
	}

	dataKey.ConfirmedDeliveryDateFrom = sdc.InputParameters.ConfirmedDeliveryDateFrom
	dataKey.ConfirmedDeliveryDateTo = sdc.InputParameters.ConfirmedDeliveryDateTo

	args = append(args, dataKey.ConfirmedDeliveryDateFrom, dataKey.ConfirmedDeliveryDateTo, dataKey.ItemScheduleLineDeliveryBlockStatus, dataKey.OpenConfirmedQuantityInBaseUnit)

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, ScheduleLine, RequestedDeliveryDate, ConfirmedDeliveryDate, OrderQuantityInBaseUnit,
		ConfirmedOrderQuantityByPDTAvailCheck, OpenConfirmedQuantityInBaseUnit, StockIsFullyConfirmed, ItemScheduleLineDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` )
		AND ConfirmedDeliveryDate BETWEEN ? AND ?
		AND ItemScheduleLineDeliveryBlockStatus = ?
		AND OpenConfirmedQuantityInBaseUnit <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersItemScheduleLine(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) OrdersItemScheduleLineInIndividualProcess(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
) ([]*api_processing_data_formatter.OrdersItemScheduleLine, error) {
	args := make([]interface{}, 0)

	dataKey := psdc.ConvertToOrdersItemScheduleLineKey()

	orderItem := psdc.OrderItem

	for i := range orderItem {
		dataKey.OrderID = append(dataKey.OrderID, (orderItem)[i].OrderID)
		dataKey.OrderItem = append(dataKey.OrderItem, (orderItem)[i].OrderItem)
	}

	repeat := strings.Repeat("(?,?),", len(dataKey.OrderID)-1) + "(?,?)"
	for i := range dataKey.OrderID {
		args = append(args, dataKey.OrderID[i], dataKey.OrderItem[i])
	}

	args = append(args, dataKey.ItemScheduleLineDeliveryBlockStatus, dataKey.OpenConfirmedQuantityInBaseUnit)

	rows, err := f.db.Query(
		`SELECT OrderID, OrderItem, ScheduleLine, RequestedDeliveryDate, ConfirmedDeliveryDate, OrderQuantityInBaseUnit,
		ConfirmedOrderQuantityByPDTAvailCheck, OpenConfirmedQuantityInBaseUnit, StockIsFullyConfirmed, ItemScheduleLineDeliveryBlockStatus
		FROM DataPlatformMastersAndTransactionsMysqlKube.data_platform_orders_item_schedule_line_data
		WHERE (OrderID, OrderItem) IN ( `+repeat+` )
		AND ItemScheduleLineDeliveryBlockStatus = ?
		AND OpenConfirmedQuantityInBaseUnit <> ?;`, args...,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	data, err := psdc.ConvertToOrdersItemScheduleLine(rows)
	if err != nil {
		return nil, err
	}

	return data, err
}

func (f *SubFunction) CreateSdc(
	sdc *api_input_reader.SDC,
	psdc *api_processing_data_formatter.SDC,
	osdc *dpfm_api_output_formatter.SDC,
) error {
	var err error
	var e error

	wg := sync.WaitGroup{}

	psdc.MetaData = f.MetaData(sdc, psdc)
	psdc.ProcessType = f.ProcessType(sdc, psdc)

	processType := psdc.ProcessType

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		if processType.BulkProcess {
			// I-1. OrderItemの絞り込み
			psdc.OrderItem, e = f.OrderItemInBulkProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// I-2. オーダー明細納入日程行の絞り込みと取得  //I-1
			psdc.OrdersItemScheduleLine, e = f.OrdersItemScheduleLineInBulkProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		} else if processType.IndividualProcess {
			// II-1-1. OrderIDが未入出荷であり、かつ、OrderIDに入出荷伝票未登録残がある、明細の取得
			psdc.OrderItem, e = f.OrderItemInIndividualProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// II-1-2. オーダー明細納入日程行の絞り込みと取得  //II-1-2
			psdc.OrdersItemScheduleLine, e = f.OrdersItemScheduleLineInIndividualProcess(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}

		// 1-1. オーダー参照レコード・値の取得（オーダーヘッダ）  //IまたはII
		psdc.OrdersHeader, e = f.OrdersHeader(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		// 1-2-0. DeliveryDocumentItem  //(I-2またはII-1-2)
		psdc.DeliveryDocumentItem = f.DeliveryDocumentItem(sdc, psdc)

		// 1-2-1. オーダー参照レコード・値の取得（オーダー明細）  //IまたはII
		psdc.OrdersItem, e = f.OrdersItem(sdc, psdc)
		if e != nil {
			err = e
			return
		}

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 1-3. DeliveryDocument  //IまたはII
			psdc.CalculateDeliveryDocument, e = f.CalculateDeliveryDocument(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 1-5. InvoiceDocumentDate  //1-2-1
			psdc.InvoiceDocumentDate, e = f.InvoiceDocumentDate(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 4-1. HeaderGrossWeight  //1-2-1
			psdc.HeaderGrossWeight = f.HeaderGrossWeight(sdc, psdc)

			// 4-2. HeaderNetWeight  //1-2-1
			psdc.HeaderNetWeight = f.HeaderNetWeight(sdc, psdc)
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 1-40. オーダー参照レコード・値の取得（オーダーパートナ）  //1-1
			psdc.Partner, e = f.Partner(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}(wg)

		wg.Add(1)
		go func(wg *sync.WaitGroup) {
			defer wg.Done()
			// 6-1. Orders Address からの住所データの取得  //IまたはII
			psdc.Address, e = f.Address(sdc, psdc)
			if e != nil {
				err = e
				return
			}

			// 6-2. AddressIDの登録(ユーザーが任意の住所を入力ファイルで指定した場合)
			psdc.Address, e = f.AddressFromInput(sdc, psdc)
			if e != nil {
				err = e
				return
			}
		}(wg)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 1-4. DocumentDate
		psdc.DocumentDate = f.DocumentDate(sdc, psdc)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 99-1-1. CreationDate(Header)
		psdc.CreationDateHeader = f.CreationDateHeader(sdc, psdc)

		// 99-2-1. LastChangeDate(Header)
		psdc.LastChangeDateHeader = f.LastChangeDateHeader(sdc, psdc)

		// 99-3-1. CreationTime(Header)
		psdc.CreationTimeHeader = f.CreationTimeHeader(sdc, psdc)

		// 99-4-1. LastChangeTime(Header)
		psdc.LastChangeTimeHeader = f.LastChangeTimeHeader(sdc, psdc)
	}(&wg)

	wg.Add(1)
	go func(wg *sync.WaitGroup) {
		defer wg.Done()
		// 99-1-2. CreationDate(Item)
		psdc.CreationDateItem = f.CreationDateItem(sdc, psdc)

		// 99-2-2. LastChangeDate(Item)
		psdc.LastChangeDateItem = f.LastChangeDateItem(sdc, psdc)

		// 99-3-2. CreationTime(Item)
		psdc.CreationTimeItem = f.CreationTimeItem(sdc, psdc)

		// 99-4-2. LastChangeTime(Item)
		psdc.LastChangeTimeItem = f.LastChangeTimeItem(sdc, psdc)
	}(&wg)

	wg.Wait()
	if err != nil {
		return err
	}

	f.l.Info(psdc)

	err = f.SetValue(sdc, osdc, psdc)
	if err != nil {
		return err
	}

	return nil
}
