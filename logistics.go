package temu

import (
	"context"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/hiscaler/temu-go/entity"
	"github.com/hiscaler/temu-go/normal"
	"github.com/hiscaler/temu-go/validators/is"
	"gopkg.in/guregu/null.v4"
)

// 物流服务

type logisticsService service

// Companies 查询发货快递公司
// https://seller.kuajingmaihuo.com/sop/view/889973754324016047#wjtGTK
func (s logisticsService) Companies(ctx context.Context) (items []entity.LogisticsCompany, err error) {
	var result = struct {
		normal.Response
		Result struct {
			ShipList []entity.LogisticsCompany `json:"shipList"` // 快递公司列表
		} `json:"result"`
	}{}
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetResult(&result).
		Post("bg.logistics.company.get")
	if err = recheckError(resp, result.Response, err); err != nil {
		return
	}

	return result.Result.ShipList, nil
}

// Company 根据 ID 查询发货快递公司
func (s logisticsService) Company(ctx context.Context, shipId int) (item entity.LogisticsCompany, err error) {
	items, err := s.Companies(ctx)
	if err != nil {
		return
	}

	for _, company := range items {
		if company.ShipId == shipId {
			return company, nil
		}
	}
	return item, ErrNotFound
}

// 平台推荐物流商匹配接口
// https://seller.kuajingmaihuo.com/sop/view/889973754324016047#16WiXI

type LogisticsMatchRequest struct {
	DeliveryAddressId         int64                 `json:"deliveryAddressId,omitempty"`   // 发货地址
	PredictTotalPackageWeight int                   `json:"predictTotalPackageWeight"`     // 预估总包裹重量，单位g
	UrgencyType               null.Int              `json:"urgencyType,omitempty"`         // 是否是紧急发货单，0-普通 1-急采
	SubWarehouseId            int64                 `json:"subWarehouseId"`                // 收货子仓 ID
	QueryStandbyExpress       null.Bool             `json:"queryStandbyExpress,omitempty"` // 是否查询备用快递服务商, false-不查询 true-查询
	TotalPackageNum           int                   `json:"totalPackageNum"`               // 包裹件数
	ReceiveAddressInfo        entity.ReceiveAddress `json:"receiveAddressInfo,omitempty"`  // 收货地址
	DeliveryOrderSns          []string              `json:"deliveryOrderSns,omitempty"`    // 发货单列表
}

func (m LogisticsMatchRequest) Validate() error {
	return validation.ValidateStruct(&m,
		validation.Field(&m.PredictTotalPackageWeight,
			validation.Required.Error("预估总包裹重量不能为空"),
			validation.Min(1).Error("预估总包裹重量不能小于 {.min}"),
		),
		validation.Field(&m.TotalPackageNum,
			validation.Required.Error("包裹件数不能为空"),
			validation.Min(1).Error("包裹件数不能小于 {.min}"),
		),
		validation.Field(&m.SubWarehouseId, validation.Required.Error("收货子仓不能为空")),
		validation.Field(&m.DeliveryOrderSns,
			validation.Required.Error("发货单列表不能为空"),
			validation.Each(validation.By(is.ShipOrderNumber())),
		),
		validation.Field(&m.ReceiveAddressInfo, validation.Required.Error("收货地址不能为空")),
	)
}

func (s logisticsService) Match(ctx context.Context, request LogisticsMatchRequest) (items []entity.LogisticsMatch, err error) {
	if err = request.Validate(); err != nil {
		return
	}

	var result = struct {
		normal.Response
		Result []entity.LogisticsMatch `json:"result"`
	}{}
	resp, err := s.httpClient.R().
		SetContext(ctx).
		SetBody(request).
		SetResult(&result).
		Post("bg.shiporderv2.logisticsmatch.get")
	if err = recheckError(resp, result.Response, err); err != nil {
		return
	}

	return result.Result, nil
}
