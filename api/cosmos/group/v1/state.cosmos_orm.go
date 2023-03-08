// Code generated by protoc-gen-go-cosmos-orm. DO NOT EDIT.

package groupv1

import (
	context "context"
	ormlist "github.com/cosmos/cosmos-sdk/orm/model/ormlist"
	ormtable "github.com/cosmos/cosmos-sdk/orm/model/ormtable"
	ormerrors "github.com/cosmos/cosmos-sdk/orm/types/ormerrors"
	timestamppb "google.golang.org/protobuf/types/known/timestamppb"
)

type GroupInfoTable interface {
	Insert(ctx context.Context, groupInfo *GroupInfo) error
	InsertReturningID(ctx context.Context, groupInfo *GroupInfo) (uint64, error)
	Update(ctx context.Context, groupInfo *GroupInfo) error
	Save(ctx context.Context, groupInfo *GroupInfo) error
	Delete(ctx context.Context, groupInfo *GroupInfo) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*GroupInfo, error)
	List(ctx context.Context, prefixKey GroupInfoIndexKey, opts ...ormlist.Option) (GroupInfoIterator, error)
	ListRange(ctx context.Context, from, to GroupInfoIndexKey, opts ...ormlist.Option) (GroupInfoIterator, error)
	DeleteBy(ctx context.Context, prefixKey GroupInfoIndexKey) error
	DeleteRange(ctx context.Context, from, to GroupInfoIndexKey) error

	doNotImplement()
}

type GroupInfoIterator struct {
	ormtable.Iterator
}

func (i GroupInfoIterator) Value() (*GroupInfo, error) {
	var groupInfo GroupInfo
	err := i.UnmarshalMessage(&groupInfo)
	return &groupInfo, err
}

type GroupInfoIndexKey interface {
	id() uint32
	values() []interface{}
	groupInfoIndexKey()
}

// primary key starting index..
type GroupInfoPrimaryKey = GroupInfoIdIndexKey

type GroupInfoIdIndexKey struct {
	vs []interface{}
}

func (x GroupInfoIdIndexKey) id() uint32            { return 0 }
func (x GroupInfoIdIndexKey) values() []interface{} { return x.vs }
func (x GroupInfoIdIndexKey) groupInfoIndexKey()    {}

func (this GroupInfoIdIndexKey) WithId(id uint64) GroupInfoIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type GroupInfoAdminIndexKey struct {
	vs []interface{}
}

func (x GroupInfoAdminIndexKey) id() uint32            { return 1 }
func (x GroupInfoAdminIndexKey) values() []interface{} { return x.vs }
func (x GroupInfoAdminIndexKey) groupInfoIndexKey()    {}

func (this GroupInfoAdminIndexKey) WithAdmin(admin string) GroupInfoAdminIndexKey {
	this.vs = []interface{}{admin}
	return this
}

type groupInfoTable struct {
	table ormtable.AutoIncrementTable
}

func (this groupInfoTable) Insert(ctx context.Context, groupInfo *GroupInfo) error {
	return this.table.Insert(ctx, groupInfo)
}

func (this groupInfoTable) Update(ctx context.Context, groupInfo *GroupInfo) error {
	return this.table.Update(ctx, groupInfo)
}

func (this groupInfoTable) Save(ctx context.Context, groupInfo *GroupInfo) error {
	return this.table.Save(ctx, groupInfo)
}

func (this groupInfoTable) Delete(ctx context.Context, groupInfo *GroupInfo) error {
	return this.table.Delete(ctx, groupInfo)
}

func (this groupInfoTable) InsertReturningID(ctx context.Context, groupInfo *GroupInfo) (uint64, error) {
	return this.table.InsertReturningID(ctx, groupInfo)
}

func (this groupInfoTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this groupInfoTable) Get(ctx context.Context, id uint64) (*GroupInfo, error) {
	var groupInfo GroupInfo
	found, err := this.table.PrimaryKey().Get(ctx, &groupInfo, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &groupInfo, nil
}

func (this groupInfoTable) List(ctx context.Context, prefixKey GroupInfoIndexKey, opts ...ormlist.Option) (GroupInfoIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return GroupInfoIterator{it}, err
}

func (this groupInfoTable) ListRange(ctx context.Context, from, to GroupInfoIndexKey, opts ...ormlist.Option) (GroupInfoIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return GroupInfoIterator{it}, err
}

func (this groupInfoTable) DeleteBy(ctx context.Context, prefixKey GroupInfoIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this groupInfoTable) DeleteRange(ctx context.Context, from, to GroupInfoIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this groupInfoTable) doNotImplement() {}

var _ GroupInfoTable = groupInfoTable{}

func NewGroupInfoTable(db ormtable.Schema) (GroupInfoTable, error) {
	table := db.GetTable(&GroupInfo{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&GroupInfo{}).ProtoReflect().Descriptor().FullName()))
	}
	return groupInfoTable{table.(ormtable.AutoIncrementTable)}, nil
}

type GroupMemberTable interface {
	Insert(ctx context.Context, groupMember *GroupMember) error
	Update(ctx context.Context, groupMember *GroupMember) error
	Save(ctx context.Context, groupMember *GroupMember) error
	Delete(ctx context.Context, groupMember *GroupMember) error
	Has(ctx context.Context, group_id uint64, member_address string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, group_id uint64, member_address string) (*GroupMember, error)
	List(ctx context.Context, prefixKey GroupMemberIndexKey, opts ...ormlist.Option) (GroupMemberIterator, error)
	ListRange(ctx context.Context, from, to GroupMemberIndexKey, opts ...ormlist.Option) (GroupMemberIterator, error)
	DeleteBy(ctx context.Context, prefixKey GroupMemberIndexKey) error
	DeleteRange(ctx context.Context, from, to GroupMemberIndexKey) error

	doNotImplement()
}

type GroupMemberIterator struct {
	ormtable.Iterator
}

func (i GroupMemberIterator) Value() (*GroupMember, error) {
	var groupMember GroupMember
	err := i.UnmarshalMessage(&groupMember)
	return &groupMember, err
}

type GroupMemberIndexKey interface {
	id() uint32
	values() []interface{}
	groupMemberIndexKey()
}

// primary key starting index..
type GroupMemberPrimaryKey = GroupMemberGroupIdMemberAddressIndexKey

type GroupMemberGroupIdMemberAddressIndexKey struct {
	vs []interface{}
}

func (x GroupMemberGroupIdMemberAddressIndexKey) id() uint32            { return 0 }
func (x GroupMemberGroupIdMemberAddressIndexKey) values() []interface{} { return x.vs }
func (x GroupMemberGroupIdMemberAddressIndexKey) groupMemberIndexKey()  {}

func (this GroupMemberGroupIdMemberAddressIndexKey) WithGroupId(group_id uint64) GroupMemberGroupIdMemberAddressIndexKey {
	this.vs = []interface{}{group_id}
	return this
}

func (this GroupMemberGroupIdMemberAddressIndexKey) WithGroupIdMemberAddress(group_id uint64, member_address string) GroupMemberGroupIdMemberAddressIndexKey {
	this.vs = []interface{}{group_id, member_address}
	return this
}

type GroupMemberMemberAddressIndexKey struct {
	vs []interface{}
}

func (x GroupMemberMemberAddressIndexKey) id() uint32            { return 1 }
func (x GroupMemberMemberAddressIndexKey) values() []interface{} { return x.vs }
func (x GroupMemberMemberAddressIndexKey) groupMemberIndexKey()  {}

func (this GroupMemberMemberAddressIndexKey) WithMemberAddress(member_address string) GroupMemberMemberAddressIndexKey {
	this.vs = []interface{}{member_address}
	return this
}

type groupMemberTable struct {
	table ormtable.Table
}

func (this groupMemberTable) Insert(ctx context.Context, groupMember *GroupMember) error {
	return this.table.Insert(ctx, groupMember)
}

func (this groupMemberTable) Update(ctx context.Context, groupMember *GroupMember) error {
	return this.table.Update(ctx, groupMember)
}

func (this groupMemberTable) Save(ctx context.Context, groupMember *GroupMember) error {
	return this.table.Save(ctx, groupMember)
}

func (this groupMemberTable) Delete(ctx context.Context, groupMember *GroupMember) error {
	return this.table.Delete(ctx, groupMember)
}

func (this groupMemberTable) Has(ctx context.Context, group_id uint64, member_address string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, group_id, member_address)
}

func (this groupMemberTable) Get(ctx context.Context, group_id uint64, member_address string) (*GroupMember, error) {
	var groupMember GroupMember
	found, err := this.table.PrimaryKey().Get(ctx, &groupMember, group_id, member_address)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &groupMember, nil
}

func (this groupMemberTable) List(ctx context.Context, prefixKey GroupMemberIndexKey, opts ...ormlist.Option) (GroupMemberIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return GroupMemberIterator{it}, err
}

func (this groupMemberTable) ListRange(ctx context.Context, from, to GroupMemberIndexKey, opts ...ormlist.Option) (GroupMemberIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return GroupMemberIterator{it}, err
}

func (this groupMemberTable) DeleteBy(ctx context.Context, prefixKey GroupMemberIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this groupMemberTable) DeleteRange(ctx context.Context, from, to GroupMemberIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this groupMemberTable) doNotImplement() {}

var _ GroupMemberTable = groupMemberTable{}

func NewGroupMemberTable(db ormtable.Schema) (GroupMemberTable, error) {
	table := db.GetTable(&GroupMember{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&GroupMember{}).ProtoReflect().Descriptor().FullName()))
	}
	return groupMemberTable{table}, nil
}

type GroupPolicyInfoTable interface {
	Insert(ctx context.Context, groupPolicyInfo *GroupPolicyInfo) error
	Update(ctx context.Context, groupPolicyInfo *GroupPolicyInfo) error
	Save(ctx context.Context, groupPolicyInfo *GroupPolicyInfo) error
	Delete(ctx context.Context, groupPolicyInfo *GroupPolicyInfo) error
	Has(ctx context.Context, address string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, address string) (*GroupPolicyInfo, error)
	List(ctx context.Context, prefixKey GroupPolicyInfoIndexKey, opts ...ormlist.Option) (GroupPolicyInfoIterator, error)
	ListRange(ctx context.Context, from, to GroupPolicyInfoIndexKey, opts ...ormlist.Option) (GroupPolicyInfoIterator, error)
	DeleteBy(ctx context.Context, prefixKey GroupPolicyInfoIndexKey) error
	DeleteRange(ctx context.Context, from, to GroupPolicyInfoIndexKey) error

	doNotImplement()
}

type GroupPolicyInfoIterator struct {
	ormtable.Iterator
}

func (i GroupPolicyInfoIterator) Value() (*GroupPolicyInfo, error) {
	var groupPolicyInfo GroupPolicyInfo
	err := i.UnmarshalMessage(&groupPolicyInfo)
	return &groupPolicyInfo, err
}

type GroupPolicyInfoIndexKey interface {
	id() uint32
	values() []interface{}
	groupPolicyInfoIndexKey()
}

// primary key starting index..
type GroupPolicyInfoPrimaryKey = GroupPolicyInfoAddressIndexKey

type GroupPolicyInfoAddressIndexKey struct {
	vs []interface{}
}

func (x GroupPolicyInfoAddressIndexKey) id() uint32               { return 0 }
func (x GroupPolicyInfoAddressIndexKey) values() []interface{}    { return x.vs }
func (x GroupPolicyInfoAddressIndexKey) groupPolicyInfoIndexKey() {}

func (this GroupPolicyInfoAddressIndexKey) WithAddress(address string) GroupPolicyInfoAddressIndexKey {
	this.vs = []interface{}{address}
	return this
}

type GroupPolicyInfoGroupIdIndexKey struct {
	vs []interface{}
}

func (x GroupPolicyInfoGroupIdIndexKey) id() uint32               { return 1 }
func (x GroupPolicyInfoGroupIdIndexKey) values() []interface{}    { return x.vs }
func (x GroupPolicyInfoGroupIdIndexKey) groupPolicyInfoIndexKey() {}

func (this GroupPolicyInfoGroupIdIndexKey) WithGroupId(group_id uint64) GroupPolicyInfoGroupIdIndexKey {
	this.vs = []interface{}{group_id}
	return this
}

type GroupPolicyInfoAdminIndexKey struct {
	vs []interface{}
}

func (x GroupPolicyInfoAdminIndexKey) id() uint32               { return 2 }
func (x GroupPolicyInfoAdminIndexKey) values() []interface{}    { return x.vs }
func (x GroupPolicyInfoAdminIndexKey) groupPolicyInfoIndexKey() {}

func (this GroupPolicyInfoAdminIndexKey) WithAdmin(admin string) GroupPolicyInfoAdminIndexKey {
	this.vs = []interface{}{admin}
	return this
}

type groupPolicyInfoTable struct {
	table ormtable.Table
}

func (this groupPolicyInfoTable) Insert(ctx context.Context, groupPolicyInfo *GroupPolicyInfo) error {
	return this.table.Insert(ctx, groupPolicyInfo)
}

func (this groupPolicyInfoTable) Update(ctx context.Context, groupPolicyInfo *GroupPolicyInfo) error {
	return this.table.Update(ctx, groupPolicyInfo)
}

func (this groupPolicyInfoTable) Save(ctx context.Context, groupPolicyInfo *GroupPolicyInfo) error {
	return this.table.Save(ctx, groupPolicyInfo)
}

func (this groupPolicyInfoTable) Delete(ctx context.Context, groupPolicyInfo *GroupPolicyInfo) error {
	return this.table.Delete(ctx, groupPolicyInfo)
}

func (this groupPolicyInfoTable) Has(ctx context.Context, address string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, address)
}

func (this groupPolicyInfoTable) Get(ctx context.Context, address string) (*GroupPolicyInfo, error) {
	var groupPolicyInfo GroupPolicyInfo
	found, err := this.table.PrimaryKey().Get(ctx, &groupPolicyInfo, address)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &groupPolicyInfo, nil
}

func (this groupPolicyInfoTable) List(ctx context.Context, prefixKey GroupPolicyInfoIndexKey, opts ...ormlist.Option) (GroupPolicyInfoIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return GroupPolicyInfoIterator{it}, err
}

func (this groupPolicyInfoTable) ListRange(ctx context.Context, from, to GroupPolicyInfoIndexKey, opts ...ormlist.Option) (GroupPolicyInfoIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return GroupPolicyInfoIterator{it}, err
}

func (this groupPolicyInfoTable) DeleteBy(ctx context.Context, prefixKey GroupPolicyInfoIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this groupPolicyInfoTable) DeleteRange(ctx context.Context, from, to GroupPolicyInfoIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this groupPolicyInfoTable) doNotImplement() {}

var _ GroupPolicyInfoTable = groupPolicyInfoTable{}

func NewGroupPolicyInfoTable(db ormtable.Schema) (GroupPolicyInfoTable, error) {
	table := db.GetTable(&GroupPolicyInfo{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&GroupPolicyInfo{}).ProtoReflect().Descriptor().FullName()))
	}
	return groupPolicyInfoTable{table}, nil
}

type ProposalTable interface {
	Insert(ctx context.Context, proposal *Proposal) error
	InsertReturningID(ctx context.Context, proposal *Proposal) (uint64, error)
	Update(ctx context.Context, proposal *Proposal) error
	Save(ctx context.Context, proposal *Proposal) error
	Delete(ctx context.Context, proposal *Proposal) error
	Has(ctx context.Context, id uint64) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, id uint64) (*Proposal, error)
	List(ctx context.Context, prefixKey ProposalIndexKey, opts ...ormlist.Option) (ProposalIterator, error)
	ListRange(ctx context.Context, from, to ProposalIndexKey, opts ...ormlist.Option) (ProposalIterator, error)
	DeleteBy(ctx context.Context, prefixKey ProposalIndexKey) error
	DeleteRange(ctx context.Context, from, to ProposalIndexKey) error

	doNotImplement()
}

type ProposalIterator struct {
	ormtable.Iterator
}

func (i ProposalIterator) Value() (*Proposal, error) {
	var proposal Proposal
	err := i.UnmarshalMessage(&proposal)
	return &proposal, err
}

type ProposalIndexKey interface {
	id() uint32
	values() []interface{}
	proposalIndexKey()
}

// primary key starting index..
type ProposalPrimaryKey = ProposalIdIndexKey

type ProposalIdIndexKey struct {
	vs []interface{}
}

func (x ProposalIdIndexKey) id() uint32            { return 0 }
func (x ProposalIdIndexKey) values() []interface{} { return x.vs }
func (x ProposalIdIndexKey) proposalIndexKey()     {}

func (this ProposalIdIndexKey) WithId(id uint64) ProposalIdIndexKey {
	this.vs = []interface{}{id}
	return this
}

type ProposalGroupPolicyAddressIndexKey struct {
	vs []interface{}
}

func (x ProposalGroupPolicyAddressIndexKey) id() uint32            { return 1 }
func (x ProposalGroupPolicyAddressIndexKey) values() []interface{} { return x.vs }
func (x ProposalGroupPolicyAddressIndexKey) proposalIndexKey()     {}

func (this ProposalGroupPolicyAddressIndexKey) WithGroupPolicyAddress(group_policy_address string) ProposalGroupPolicyAddressIndexKey {
	this.vs = []interface{}{group_policy_address}
	return this
}

type ProposalVotingPeriodEndIndexKey struct {
	vs []interface{}
}

func (x ProposalVotingPeriodEndIndexKey) id() uint32            { return 2 }
func (x ProposalVotingPeriodEndIndexKey) values() []interface{} { return x.vs }
func (x ProposalVotingPeriodEndIndexKey) proposalIndexKey()     {}

func (this ProposalVotingPeriodEndIndexKey) WithVotingPeriodEnd(voting_period_end *timestamppb.Timestamp) ProposalVotingPeriodEndIndexKey {
	this.vs = []interface{}{voting_period_end}
	return this
}

type proposalTable struct {
	table ormtable.AutoIncrementTable
}

func (this proposalTable) Insert(ctx context.Context, proposal *Proposal) error {
	return this.table.Insert(ctx, proposal)
}

func (this proposalTable) Update(ctx context.Context, proposal *Proposal) error {
	return this.table.Update(ctx, proposal)
}

func (this proposalTable) Save(ctx context.Context, proposal *Proposal) error {
	return this.table.Save(ctx, proposal)
}

func (this proposalTable) Delete(ctx context.Context, proposal *Proposal) error {
	return this.table.Delete(ctx, proposal)
}

func (this proposalTable) InsertReturningID(ctx context.Context, proposal *Proposal) (uint64, error) {
	return this.table.InsertReturningID(ctx, proposal)
}

func (this proposalTable) Has(ctx context.Context, id uint64) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, id)
}

func (this proposalTable) Get(ctx context.Context, id uint64) (*Proposal, error) {
	var proposal Proposal
	found, err := this.table.PrimaryKey().Get(ctx, &proposal, id)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &proposal, nil
}

func (this proposalTable) List(ctx context.Context, prefixKey ProposalIndexKey, opts ...ormlist.Option) (ProposalIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return ProposalIterator{it}, err
}

func (this proposalTable) ListRange(ctx context.Context, from, to ProposalIndexKey, opts ...ormlist.Option) (ProposalIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return ProposalIterator{it}, err
}

func (this proposalTable) DeleteBy(ctx context.Context, prefixKey ProposalIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this proposalTable) DeleteRange(ctx context.Context, from, to ProposalIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this proposalTable) doNotImplement() {}

var _ ProposalTable = proposalTable{}

func NewProposalTable(db ormtable.Schema) (ProposalTable, error) {
	table := db.GetTable(&Proposal{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Proposal{}).ProtoReflect().Descriptor().FullName()))
	}
	return proposalTable{table.(ormtable.AutoIncrementTable)}, nil
}

type VoteTable interface {
	Insert(ctx context.Context, vote *Vote) error
	Update(ctx context.Context, vote *Vote) error
	Save(ctx context.Context, vote *Vote) error
	Delete(ctx context.Context, vote *Vote) error
	Has(ctx context.Context, proposal_id uint64, voter string) (found bool, err error)
	// Get returns nil and an error which responds true to ormerrors.IsNotFound() if the record was not found.
	Get(ctx context.Context, proposal_id uint64, voter string) (*Vote, error)
	List(ctx context.Context, prefixKey VoteIndexKey, opts ...ormlist.Option) (VoteIterator, error)
	ListRange(ctx context.Context, from, to VoteIndexKey, opts ...ormlist.Option) (VoteIterator, error)
	DeleteBy(ctx context.Context, prefixKey VoteIndexKey) error
	DeleteRange(ctx context.Context, from, to VoteIndexKey) error

	doNotImplement()
}

type VoteIterator struct {
	ormtable.Iterator
}

func (i VoteIterator) Value() (*Vote, error) {
	var vote Vote
	err := i.UnmarshalMessage(&vote)
	return &vote, err
}

type VoteIndexKey interface {
	id() uint32
	values() []interface{}
	voteIndexKey()
}

// primary key starting index..
type VotePrimaryKey = VoteProposalIdVoterIndexKey

type VoteProposalIdVoterIndexKey struct {
	vs []interface{}
}

func (x VoteProposalIdVoterIndexKey) id() uint32            { return 0 }
func (x VoteProposalIdVoterIndexKey) values() []interface{} { return x.vs }
func (x VoteProposalIdVoterIndexKey) voteIndexKey()         {}

func (this VoteProposalIdVoterIndexKey) WithProposalId(proposal_id uint64) VoteProposalIdVoterIndexKey {
	this.vs = []interface{}{proposal_id}
	return this
}

func (this VoteProposalIdVoterIndexKey) WithProposalIdVoter(proposal_id uint64, voter string) VoteProposalIdVoterIndexKey {
	this.vs = []interface{}{proposal_id, voter}
	return this
}

type VoteVoterIndexKey struct {
	vs []interface{}
}

func (x VoteVoterIndexKey) id() uint32            { return 1 }
func (x VoteVoterIndexKey) values() []interface{} { return x.vs }
func (x VoteVoterIndexKey) voteIndexKey()         {}

func (this VoteVoterIndexKey) WithVoter(voter string) VoteVoterIndexKey {
	this.vs = []interface{}{voter}
	return this
}

type voteTable struct {
	table ormtable.Table
}

func (this voteTable) Insert(ctx context.Context, vote *Vote) error {
	return this.table.Insert(ctx, vote)
}

func (this voteTable) Update(ctx context.Context, vote *Vote) error {
	return this.table.Update(ctx, vote)
}

func (this voteTable) Save(ctx context.Context, vote *Vote) error {
	return this.table.Save(ctx, vote)
}

func (this voteTable) Delete(ctx context.Context, vote *Vote) error {
	return this.table.Delete(ctx, vote)
}

func (this voteTable) Has(ctx context.Context, proposal_id uint64, voter string) (found bool, err error) {
	return this.table.PrimaryKey().Has(ctx, proposal_id, voter)
}

func (this voteTable) Get(ctx context.Context, proposal_id uint64, voter string) (*Vote, error) {
	var vote Vote
	found, err := this.table.PrimaryKey().Get(ctx, &vote, proposal_id, voter)
	if err != nil {
		return nil, err
	}
	if !found {
		return nil, ormerrors.NotFound
	}
	return &vote, nil
}

func (this voteTable) List(ctx context.Context, prefixKey VoteIndexKey, opts ...ormlist.Option) (VoteIterator, error) {
	it, err := this.table.GetIndexByID(prefixKey.id()).List(ctx, prefixKey.values(), opts...)
	return VoteIterator{it}, err
}

func (this voteTable) ListRange(ctx context.Context, from, to VoteIndexKey, opts ...ormlist.Option) (VoteIterator, error) {
	it, err := this.table.GetIndexByID(from.id()).ListRange(ctx, from.values(), to.values(), opts...)
	return VoteIterator{it}, err
}

func (this voteTable) DeleteBy(ctx context.Context, prefixKey VoteIndexKey) error {
	return this.table.GetIndexByID(prefixKey.id()).DeleteBy(ctx, prefixKey.values()...)
}

func (this voteTable) DeleteRange(ctx context.Context, from, to VoteIndexKey) error {
	return this.table.GetIndexByID(from.id()).DeleteRange(ctx, from.values(), to.values())
}

func (this voteTable) doNotImplement() {}

var _ VoteTable = voteTable{}

func NewVoteTable(db ormtable.Schema) (VoteTable, error) {
	table := db.GetTable(&Vote{})
	if table == nil {
		return nil, ormerrors.TableNotFound.Wrap(string((&Vote{}).ProtoReflect().Descriptor().FullName()))
	}
	return voteTable{table}, nil
}

type StateStore interface {
	GroupInfoTable() GroupInfoTable
	GroupMemberTable() GroupMemberTable
	GroupPolicyInfoTable() GroupPolicyInfoTable
	ProposalTable() ProposalTable
	VoteTable() VoteTable

	doNotImplement()
}

type stateStore struct {
	groupInfo       GroupInfoTable
	groupMember     GroupMemberTable
	groupPolicyInfo GroupPolicyInfoTable
	proposal        ProposalTable
	vote            VoteTable
}

func (x stateStore) GroupInfoTable() GroupInfoTable {
	return x.groupInfo
}

func (x stateStore) GroupMemberTable() GroupMemberTable {
	return x.groupMember
}

func (x stateStore) GroupPolicyInfoTable() GroupPolicyInfoTable {
	return x.groupPolicyInfo
}

func (x stateStore) ProposalTable() ProposalTable {
	return x.proposal
}

func (x stateStore) VoteTable() VoteTable {
	return x.vote
}

func (stateStore) doNotImplement() {}

var _ StateStore = stateStore{}

func NewStateStore(db ormtable.Schema) (StateStore, error) {
	groupInfoTable, err := NewGroupInfoTable(db)
	if err != nil {
		return nil, err
	}

	groupMemberTable, err := NewGroupMemberTable(db)
	if err != nil {
		return nil, err
	}

	groupPolicyInfoTable, err := NewGroupPolicyInfoTable(db)
	if err != nil {
		return nil, err
	}

	proposalTable, err := NewProposalTable(db)
	if err != nil {
		return nil, err
	}

	voteTable, err := NewVoteTable(db)
	if err != nil {
		return nil, err
	}

	return stateStore{
		groupInfoTable,
		groupMemberTable,
		groupPolicyInfoTable,
		proposalTable,
		voteTable,
	}, nil
}
