// Code generated by protoc-gen-go.
// source: game.proto
// DO NOT EDIT!

package types

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

type Game struct {
	// 默认是由创建这局游戏的txHash作为gameId
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	// create 1 -> Match 2 -> Cancel 3 -> Close 4
	Status int32 `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	// 创建时间
	CreateTime int64 `protobuf:"varint,3,opt,name=createTime" json:"createTime,omitempty"`
	// 匹配时间(何时参与对赌）
	MatchTime int64 `protobuf:"varint,4,opt,name=matchTime" json:"matchTime,omitempty"`
	// 状态close的时间（包括cancel）
	Closetime int64 `protobuf:"varint,5,opt,name=closetime" json:"closetime,omitempty"`
	// 赌注
	Value int64 `protobuf:"varint,6,opt,name=value" json:"value,omitempty"`
	// 发起者账号地址
	CreateAddress string `protobuf:"bytes,7,opt,name=createAddress" json:"createAddress,omitempty"`
	// 对赌者账号地址
	MatchAddress string `protobuf:"bytes,8,opt,name=matchAddress" json:"matchAddress,omitempty"`
	// hash 类型，预留字段
	HashType string `protobuf:"bytes,9,opt,name=hashType" json:"hashType,omitempty"`
	// 庄家创建游戏时，庄家自己出拳结果加密后的hash值
	HashValue []byte `protobuf:"bytes,10,opt,name=hashValue,proto3" json:"hashValue,omitempty"`
	// 用来公布庄家出拳结果的私钥
	Secret string `protobuf:"bytes,11,opt,name=secret" json:"secret,omitempty"`
	// 0 平局，1 庄家获胜，2 matcher获胜，3 庄家开奖超时，matcher获胜，并获得本局所有赌资
	Result int32 `protobuf:"varint,12,opt,name=result" json:"result,omitempty"`
	// matcher 出拳结果
	Guess int32 `protobuf:"varint,13,opt,name=guess" json:"guess,omitempty"`
}

func (m *Game) Reset()                    { *m = Game{} }
func (m *Game) String() string            { return proto.CompactTextString(m) }
func (*Game) ProtoMessage()               {}
func (*Game) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{0} }

func (m *Game) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *Game) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *Game) GetCreateTime() int64 {
	if m != nil {
		return m.CreateTime
	}
	return 0
}

func (m *Game) GetMatchTime() int64 {
	if m != nil {
		return m.MatchTime
	}
	return 0
}

func (m *Game) GetClosetime() int64 {
	if m != nil {
		return m.Closetime
	}
	return 0
}

func (m *Game) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *Game) GetCreateAddress() string {
	if m != nil {
		return m.CreateAddress
	}
	return ""
}

func (m *Game) GetMatchAddress() string {
	if m != nil {
		return m.MatchAddress
	}
	return ""
}

func (m *Game) GetHashType() string {
	if m != nil {
		return m.HashType
	}
	return ""
}

func (m *Game) GetHashValue() []byte {
	if m != nil {
		return m.HashValue
	}
	return nil
}

func (m *Game) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

func (m *Game) GetResult() int32 {
	if m != nil {
		return m.Result
	}
	return 0
}

func (m *Game) GetGuess() int32 {
	if m != nil {
		return m.Guess
	}
	return 0
}

// message for execs.game
type GameAction struct {
	// Types that are valid to be assigned to Value:
	//	*GameAction_Create
	//	*GameAction_Cancel
	//	*GameAction_Close
	//	*GameAction_Match
	Value isGameAction_Value `protobuf_oneof:"value"`
	Ty    int32              `protobuf:"varint,10,opt,name=ty" json:"ty,omitempty"`
}

func (m *GameAction) Reset()                    { *m = GameAction{} }
func (m *GameAction) String() string            { return proto.CompactTextString(m) }
func (*GameAction) ProtoMessage()               {}
func (*GameAction) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{1} }

type isGameAction_Value interface {
	isGameAction_Value()
}

type GameAction_Create struct {
	Create *GameCreate `protobuf:"bytes,1,opt,name=create,oneof"`
}
type GameAction_Cancel struct {
	Cancel *GameCancel `protobuf:"bytes,2,opt,name=cancel,oneof"`
}
type GameAction_Close struct {
	Close *GameClose `protobuf:"bytes,3,opt,name=close,oneof"`
}
type GameAction_Match struct {
	Match *GameMatch `protobuf:"bytes,4,opt,name=match,oneof"`
}

func (*GameAction_Create) isGameAction_Value() {}
func (*GameAction_Cancel) isGameAction_Value() {}
func (*GameAction_Close) isGameAction_Value()  {}
func (*GameAction_Match) isGameAction_Value()  {}

func (m *GameAction) GetValue() isGameAction_Value {
	if m != nil {
		return m.Value
	}
	return nil
}

func (m *GameAction) GetCreate() *GameCreate {
	if x, ok := m.GetValue().(*GameAction_Create); ok {
		return x.Create
	}
	return nil
}

func (m *GameAction) GetCancel() *GameCancel {
	if x, ok := m.GetValue().(*GameAction_Cancel); ok {
		return x.Cancel
	}
	return nil
}

func (m *GameAction) GetClose() *GameClose {
	if x, ok := m.GetValue().(*GameAction_Close); ok {
		return x.Close
	}
	return nil
}

func (m *GameAction) GetMatch() *GameMatch {
	if x, ok := m.GetValue().(*GameAction_Match); ok {
		return x.Match
	}
	return nil
}

func (m *GameAction) GetTy() int32 {
	if m != nil {
		return m.Ty
	}
	return 0
}

// XXX_OneofFuncs is for the internal use of the proto package.
func (*GameAction) XXX_OneofFuncs() (func(msg proto.Message, b *proto.Buffer) error, func(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error), func(msg proto.Message) (n int), []interface{}) {
	return _GameAction_OneofMarshaler, _GameAction_OneofUnmarshaler, _GameAction_OneofSizer, []interface{}{
		(*GameAction_Create)(nil),
		(*GameAction_Cancel)(nil),
		(*GameAction_Close)(nil),
		(*GameAction_Match)(nil),
	}
}

func _GameAction_OneofMarshaler(msg proto.Message, b *proto.Buffer) error {
	m := msg.(*GameAction)
	// value
	switch x := m.Value.(type) {
	case *GameAction_Create:
		b.EncodeVarint(1<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Create); err != nil {
			return err
		}
	case *GameAction_Cancel:
		b.EncodeVarint(2<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Cancel); err != nil {
			return err
		}
	case *GameAction_Close:
		b.EncodeVarint(3<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Close); err != nil {
			return err
		}
	case *GameAction_Match:
		b.EncodeVarint(4<<3 | proto.WireBytes)
		if err := b.EncodeMessage(x.Match); err != nil {
			return err
		}
	case nil:
	default:
		return fmt.Errorf("GameAction.Value has unexpected type %T", x)
	}
	return nil
}

func _GameAction_OneofUnmarshaler(msg proto.Message, tag, wire int, b *proto.Buffer) (bool, error) {
	m := msg.(*GameAction)
	switch tag {
	case 1: // value.create
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GameCreate)
		err := b.DecodeMessage(msg)
		m.Value = &GameAction_Create{msg}
		return true, err
	case 2: // value.cancel
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GameCancel)
		err := b.DecodeMessage(msg)
		m.Value = &GameAction_Cancel{msg}
		return true, err
	case 3: // value.close
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GameClose)
		err := b.DecodeMessage(msg)
		m.Value = &GameAction_Close{msg}
		return true, err
	case 4: // value.match
		if wire != proto.WireBytes {
			return true, proto.ErrInternalBadWireType
		}
		msg := new(GameMatch)
		err := b.DecodeMessage(msg)
		m.Value = &GameAction_Match{msg}
		return true, err
	default:
		return false, nil
	}
}

func _GameAction_OneofSizer(msg proto.Message) (n int) {
	m := msg.(*GameAction)
	// value
	switch x := m.Value.(type) {
	case *GameAction_Create:
		s := proto.Size(x.Create)
		n += proto.SizeVarint(1<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GameAction_Cancel:
		s := proto.Size(x.Cancel)
		n += proto.SizeVarint(2<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GameAction_Close:
		s := proto.Size(x.Close)
		n += proto.SizeVarint(3<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case *GameAction_Match:
		s := proto.Size(x.Match)
		n += proto.SizeVarint(4<<3 | proto.WireBytes)
		n += proto.SizeVarint(uint64(s))
		n += s
	case nil:
	default:
		panic(fmt.Sprintf("proto: unexpected type %T in oneof", x))
	}
	return n
}

type GameMatch struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	Guess  int32  `protobuf:"varint,2,opt,name=guess" json:"guess,omitempty"`
}

func (m *GameMatch) Reset()                    { *m = GameMatch{} }
func (m *GameMatch) String() string            { return proto.CompactTextString(m) }
func (*GameMatch) ProtoMessage()               {}
func (*GameMatch) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{2} }

func (m *GameMatch) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *GameMatch) GetGuess() int32 {
	if m != nil {
		return m.Guess
	}
	return 0
}

type GameCancel struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *GameCancel) Reset()                    { *m = GameCancel{} }
func (m *GameCancel) String() string            { return proto.CompactTextString(m) }
func (*GameCancel) ProtoMessage()               {}
func (*GameCancel) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{3} }

func (m *GameCancel) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type GameClose struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	Secret string `protobuf:"bytes,2,opt,name=secret" json:"secret,omitempty"`
}

func (m *GameClose) Reset()                    { *m = GameClose{} }
func (m *GameClose) String() string            { return proto.CompactTextString(m) }
func (*GameClose) ProtoMessage()               {}
func (*GameClose) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{4} }

func (m *GameClose) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *GameClose) GetSecret() string {
	if m != nil {
		return m.Secret
	}
	return ""
}

type GameCreate struct {
	Value int64 `protobuf:"varint,1,opt,name=value" json:"value,omitempty"`
	// 加密用的算法
	HashType string `protobuf:"bytes,2,opt,name=hashType" json:"hashType,omitempty"`
	// 加密后的值
	HashValue []byte `protobuf:"bytes,3,opt,name=hashValue,proto3" json:"hashValue,omitempty"`
}

func (m *GameCreate) Reset()                    { *m = GameCreate{} }
func (m *GameCreate) String() string            { return proto.CompactTextString(m) }
func (*GameCreate) ProtoMessage()               {}
func (*GameCreate) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{5} }

func (m *GameCreate) GetValue() int64 {
	if m != nil {
		return m.Value
	}
	return 0
}

func (m *GameCreate) GetHashType() string {
	if m != nil {
		return m.HashType
	}
	return ""
}

func (m *GameCreate) GetHashValue() []byte {
	if m != nil {
		return m.HashValue
	}
	return nil
}

// queryByAddr 和 queryByStatus共用同一个结构体
type QueryGameListByStatusAndAddr struct {
	// 优先根据status查询,status不可为空
	Status int32 `protobuf:"varint,1,opt,name=status" json:"status,omitempty"`
	// 二级搜索，如果要查询一个地址下的所有game信息，可以根据status，分多次查询，这样规避存储数据时的臃余情况
	Address string `protobuf:"bytes,2,opt,name=address" json:"address,omitempty"`
}

func (m *QueryGameListByStatusAndAddr) Reset()                    { *m = QueryGameListByStatusAndAddr{} }
func (m *QueryGameListByStatusAndAddr) String() string            { return proto.CompactTextString(m) }
func (*QueryGameListByStatusAndAddr) ProtoMessage()               {}
func (*QueryGameListByStatusAndAddr) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{6} }

func (m *QueryGameListByStatusAndAddr) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *QueryGameListByStatusAndAddr) GetAddress() string {
	if m != nil {
		return m.Address
	}
	return ""
}

type QueryGameInfo struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
}

func (m *QueryGameInfo) Reset()                    { *m = QueryGameInfo{} }
func (m *QueryGameInfo) String() string            { return proto.CompactTextString(m) }
func (*QueryGameInfo) ProtoMessage()               {}
func (*QueryGameInfo) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{7} }

func (m *QueryGameInfo) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

type QueryGameInfos struct {
	GameIds []string `protobuf:"bytes,1,rep,name=gameIds" json:"gameIds,omitempty"`
}

func (m *QueryGameInfos) Reset()                    { *m = QueryGameInfos{} }
func (m *QueryGameInfos) String() string            { return proto.CompactTextString(m) }
func (*QueryGameInfos) ProtoMessage()               {}
func (*QueryGameInfos) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{8} }

func (m *QueryGameInfos) GetGameIds() []string {
	if m != nil {
		return m.GameIds
	}
	return nil
}

type ReplyGameList struct {
	Games []*Game `protobuf:"bytes,1,rep,name=games" json:"games,omitempty"`
}

func (m *ReplyGameList) Reset()                    { *m = ReplyGameList{} }
func (m *ReplyGameList) String() string            { return proto.CompactTextString(m) }
func (*ReplyGameList) ProtoMessage()               {}
func (*ReplyGameList) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{9} }

func (m *ReplyGameList) GetGames() []*Game {
	if m != nil {
		return m.Games
	}
	return nil
}

type ReplyGame struct {
	Game *Game `protobuf:"bytes,1,opt,name=game" json:"game,omitempty"`
}

func (m *ReplyGame) Reset()                    { *m = ReplyGame{} }
func (m *ReplyGame) String() string            { return proto.CompactTextString(m) }
func (*ReplyGame) ProtoMessage()               {}
func (*ReplyGame) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{10} }

func (m *ReplyGame) GetGame() *Game {
	if m != nil {
		return m.Game
	}
	return nil
}

type ReceiptGame struct {
	GameId string `protobuf:"bytes,1,opt,name=gameId" json:"gameId,omitempty"`
	Status int32  `protobuf:"varint,2,opt,name=status" json:"status,omitempty"`
	// 记录上一次状态
	PrevStatus int32  `protobuf:"varint,4,opt,name=prevStatus" json:"prevStatus,omitempty"`
	Addr       string `protobuf:"bytes,5,opt,name=addr" json:"addr,omitempty"`
}

func (m *ReceiptGame) Reset()                    { *m = ReceiptGame{} }
func (m *ReceiptGame) String() string            { return proto.CompactTextString(m) }
func (*ReceiptGame) ProtoMessage()               {}
func (*ReceiptGame) Descriptor() ([]byte, []int) { return fileDescriptor8, []int{11} }

func (m *ReceiptGame) GetGameId() string {
	if m != nil {
		return m.GameId
	}
	return ""
}

func (m *ReceiptGame) GetStatus() int32 {
	if m != nil {
		return m.Status
	}
	return 0
}

func (m *ReceiptGame) GetPrevStatus() int32 {
	if m != nil {
		return m.PrevStatus
	}
	return 0
}

func (m *ReceiptGame) GetAddr() string {
	if m != nil {
		return m.Addr
	}
	return ""
}

func init() {
	proto.RegisterType((*Game)(nil), "types.Game")
	proto.RegisterType((*GameAction)(nil), "types.GameAction")
	proto.RegisterType((*GameMatch)(nil), "types.GameMatch")
	proto.RegisterType((*GameCancel)(nil), "types.GameCancel")
	proto.RegisterType((*GameClose)(nil), "types.GameClose")
	proto.RegisterType((*GameCreate)(nil), "types.GameCreate")
	proto.RegisterType((*QueryGameListByStatusAndAddr)(nil), "types.QueryGameListByStatusAndAddr")
	proto.RegisterType((*QueryGameInfo)(nil), "types.QueryGameInfo")
	proto.RegisterType((*QueryGameInfos)(nil), "types.QueryGameInfos")
	proto.RegisterType((*ReplyGameList)(nil), "types.ReplyGameList")
	proto.RegisterType((*ReplyGame)(nil), "types.ReplyGame")
	proto.RegisterType((*ReceiptGame)(nil), "types.ReceiptGame")
}

func init() { proto.RegisterFile("game.proto", fileDescriptor8) }

var fileDescriptor8 = []byte{
	// 533 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x54, 0xc1, 0x6e, 0xd3, 0x40,
	0x10, 0xad, 0x9d, 0x38, 0xa9, 0xc7, 0x49, 0x05, 0x2b, 0x84, 0x2c, 0x54, 0x41, 0x58, 0x55, 0xc2,
	0x02, 0x94, 0x43, 0x38, 0x21, 0x4e, 0x29, 0x07, 0x5a, 0x09, 0x24, 0x58, 0x2a, 0x4e, 0x5c, 0x8c,
	0x33, 0x34, 0x91, 0x9c, 0xd8, 0xd8, 0xeb, 0x4a, 0xfe, 0x4e, 0xfe, 0x84, 0x2f, 0x40, 0x33, 0xe3,
	0xba, 0x76, 0xa5, 0xe4, 0xd0, 0x9b, 0xe7, 0xbd, 0xe7, 0x9d, 0xf1, 0x9b, 0xe7, 0x05, 0xb8, 0x8e,
	0xb7, 0x38, 0xcf, 0x8b, 0xcc, 0x66, 0xca, 0xb3, 0x75, 0x8e, 0xa5, 0xfe, 0xe7, 0xc2, 0xf0, 0x53,
	0xbc, 0x45, 0xf5, 0x14, 0x46, 0xc4, 0x5e, 0xae, 0x42, 0x67, 0xe6, 0x44, 0xbe, 0x69, 0x2a, 0xc2,
	0x4b, 0x1b, 0xdb, 0xaa, 0x0c, 0xdd, 0x99, 0x13, 0x79, 0xa6, 0xa9, 0xd4, 0x73, 0x80, 0xa4, 0xc0,
	0xd8, 0xe2, 0xd5, 0x66, 0x8b, 0xe1, 0x60, 0xe6, 0x44, 0x03, 0xd3, 0x41, 0xd4, 0x29, 0xf8, 0xdb,
	0xd8, 0x26, 0x6b, 0xa6, 0x87, 0x4c, 0xdf, 0x01, 0xc4, 0x26, 0x69, 0x56, 0xa2, 0x25, 0xd6, 0x13,
	0xb6, 0x05, 0xd4, 0x13, 0xf0, 0x6e, 0xe2, 0xb4, 0xc2, 0x70, 0xc4, 0x8c, 0x14, 0xea, 0x0c, 0xa6,
	0x72, 0xfe, 0x72, 0xb5, 0x2a, 0xb0, 0x2c, 0xc3, 0x31, 0x0f, 0xda, 0x07, 0x95, 0x86, 0x09, 0xb7,
	0xb9, 0x15, 0x1d, 0xb3, 0xa8, 0x87, 0xa9, 0x67, 0x70, 0xbc, 0x8e, 0xcb, 0xf5, 0x55, 0x9d, 0x63,
	0xe8, 0x33, 0xdf, 0xd6, 0x34, 0x19, 0x3d, 0xff, 0xe0, 0xfe, 0x30, 0x73, 0xa2, 0x89, 0xb9, 0x03,
	0xd8, 0x0d, 0x4c, 0x0a, 0xb4, 0x61, 0x20, 0x2e, 0x49, 0x45, 0x78, 0x81, 0x65, 0x95, 0xda, 0x70,
	0x22, 0x2e, 0x49, 0x45, 0x5f, 0x72, 0x5d, 0xd1, 0x18, 0x53, 0x86, 0xa5, 0xd0, 0x7f, 0x1d, 0x00,
	0x32, 0x7d, 0x99, 0xd8, 0x4d, 0xb6, 0x53, 0x6f, 0x60, 0x24, 0xdf, 0xc0, 0xd6, 0x07, 0x8b, 0xc7,
	0x73, 0xde, 0xcd, 0x9c, 0x24, 0x1f, 0x99, 0xb8, 0x38, 0x32, 0x8d, 0x84, 0xc5, 0xf1, 0x2e, 0xc1,
	0x94, 0xf7, 0x71, 0x4f, 0xcc, 0x04, 0x8b, 0xf9, 0x49, 0x45, 0xe0, 0xb1, 0xab, 0xbc, 0x9f, 0x60,
	0xf1, 0xa8, 0xab, 0x25, 0xfc, 0xe2, 0xc8, 0x88, 0x80, 0x94, 0x6c, 0x11, 0xaf, 0xaa, 0xaf, 0xfc,
	0x42, 0x38, 0x29, 0x59, 0xa0, 0x4e, 0xc0, 0xb5, 0x35, 0x3b, 0xe3, 0x19, 0xd7, 0xd6, 0xe7, 0xe3,
	0x66, 0x59, 0xfa, 0x3d, 0xf8, 0xad, 0x7c, 0x6f, 0x9c, 0x5a, 0x43, 0xdc, 0xae, 0x21, 0x67, 0xe2,
	0x87, 0xcc, 0xbf, 0xef, 0x5d, 0xfd, 0x41, 0x1a, 0xf0, 0xe4, 0x07, 0xf3, 0x2a, 0x1b, 0x72, 0xbb,
	0x1b, 0xd2, 0x3f, 0x9b, 0x16, 0xe2, 0x62, 0x9b, 0x30, 0xa7, 0x9b, 0xb0, 0x6e, 0x2e, 0xdc, 0x43,
	0xb9, 0x18, 0xdc, 0xcb, 0x85, 0xfe, 0x0a, 0xa7, 0xdf, 0x2a, 0x2c, 0x6a, 0x6a, 0xf1, 0x79, 0x53,
	0xda, 0xf3, 0xfa, 0x3b, 0xff, 0x26, 0xcb, 0xdd, 0x8a, 0x42, 0xd7, 0xf9, 0x8b, 0x9c, 0xde, 0x5f,
	0x14, 0xc2, 0x38, 0x6e, 0x82, 0x2a, 0x0d, 0x6f, 0x4b, 0xfd, 0x0a, 0xa6, 0xed, 0x89, 0x97, 0xbb,
	0xdf, 0xd9, 0x5e, 0x57, 0x5e, 0xc3, 0x49, 0x4f, 0xc8, 0x87, 0x0a, 0x47, 0xdd, 0x06, 0x74, 0x68,
	0x53, 0xea, 0x05, 0x4c, 0x0d, 0xe6, 0x69, 0x3b, 0xa6, 0x7a, 0x09, 0x1e, 0x71, 0x22, 0x0c, 0x16,
	0x41, 0x67, 0xed, 0x46, 0x18, 0xfd, 0x16, 0xfc, 0xf6, 0x1d, 0xf5, 0x02, 0x86, 0x84, 0x36, 0x41,
	0xed, 0xc9, 0x99, 0xd0, 0x7f, 0x20, 0x30, 0x98, 0xe0, 0x26, 0xb7, 0x0f, 0xbd, 0x55, 0xf2, 0x02,
	0x6f, 0xc4, 0x3c, 0xce, 0xa2, 0x67, 0x3a, 0x88, 0x52, 0x30, 0x24, 0x83, 0xf8, 0xca, 0xf0, 0x0d,
	0x3f, 0xff, 0x1a, 0xf1, 0x85, 0xf6, 0xee, 0x7f, 0x00, 0x00, 0x00, 0xff, 0xff, 0x90, 0xa6, 0x53,
	0x4b, 0xde, 0x04, 0x00, 0x00,
}
