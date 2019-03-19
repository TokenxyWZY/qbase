package validator

import (
	"github.com/QOSGroup/qbase/context"
	"github.com/QOSGroup/qbase/mapper"
	"github.com/QOSGroup/qbase/types"
	abci "github.com/tendermint/tendermint/abci/types"
)

type ValidatorMapper struct {
	*mapper.BaseMapper
}

func (mapper *ValidatorMapper) ClearValidatorUpdateSet() {
	mapper.Del(ValidatorUpdateSetKey)
}

func (mapper *ValidatorMapper) GetValidatorUpdateSet() (update []abci.ValidatorUpdate) {
	mapper.Get(ValidatorUpdateSetKey, &update)
	return
}

func (mapper *ValidatorMapper) SetValidatorUpdateSet(update []abci.ValidatorUpdate) {
	mapper.Set(ValidatorUpdateSetKey, update)
}

func (mapper *ValidatorMapper) SetLastBlockProposer(address types.Address) {
	mapper.Set(LastBlockProposerKey, address)
}

func (mapper *ValidatorMapper) GetLastBlockProposer() (address types.Address, exsits bool) {
	exsits = mapper.Get(LastBlockProposerKey, &address)
	return
}

func (mapper *ValidatorMapper) IsEnableValidatorUpdated() bool {
	if v, exsits := mapper.GetBool(EnableValidatorUpdatedKey); exsits {
		return v
	}
	return false
}

func (mapper *ValidatorMapper) EnableValidatorUpdated() {
	mapper.Set(EnableValidatorUpdatedKey, true)
}

func (mapper *ValidatorMapper) DisableValidatorUpdated() {
	mapper.Set(EnableValidatorUpdatedKey, false)
}

var _ mapper.IMapper = (*ValidatorMapper)(nil)

func NewValidatorMapper() *ValidatorMapper {
	var validatorMapper = ValidatorMapper{}
	validatorMapper.BaseMapper = mapper.NewBaseMapper(nil, ValidatorMapperName)
	return &validatorMapper
}

func GetValidatorMapper(ctx context.Context) *ValidatorMapper {
	return ctx.Mapper(ValidatorMapperName).(*ValidatorMapper)
}

func (mapper *ValidatorMapper) Copy() mapper.IMapper {
	validatorMapper := &ValidatorMapper{}
	validatorMapper.BaseMapper = mapper.BaseMapper.Copy()
	return validatorMapper
}
