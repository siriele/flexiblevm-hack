// Code generated by protoc-gen-go.
// source: TangramAnalyticsEvent.proto
// DO NOT EDIT!

package generated

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// Tangram session event
type TangramSessionEvent struct {
	LastDifficulty       *int32 `protobuf:"varint,1,opt,name=last_difficulty,json=lastDifficulty" json:"last_difficulty,omitempty"`
	NumStartedTreasure   *int32 `protobuf:"varint,2,opt,name=num_started_treasure,json=numStartedTreasure" json:"num_started_treasure,omitempty"`
	NumSolvedTreasure    *int32 `protobuf:"varint,3,opt,name=num_solved_treasure,json=numSolvedTreasure" json:"num_solved_treasure,omitempty"`
	NumAbandonedTreasure *int32 `protobuf:"varint,4,opt,name=num_abandoned_treasure,json=numAbandonedTreasure" json:"num_abandoned_treasure,omitempty"`
	NumStartedClassic    *int32 `protobuf:"varint,5,opt,name=num_started_classic,json=numStartedClassic" json:"num_started_classic,omitempty"`
	NumSolvedClassic     *int32 `protobuf:"varint,6,opt,name=num_solved_classic,json=numSolvedClassic" json:"num_solved_classic,omitempty"`
	NumAbandonedClassic  *int32 `protobuf:"varint,7,opt,name=num_abandoned_classic,json=numAbandonedClassic" json:"num_abandoned_classic,omitempty"`
	NumStartedAcademy    *int32 `protobuf:"varint,8,opt,name=num_started_academy,json=numStartedAcademy" json:"num_started_academy,omitempty"`
	NumSolvedAcademy     *int32 `protobuf:"varint,9,opt,name=num_solved_academy,json=numSolvedAcademy" json:"num_solved_academy,omitempty"`
	XXX_unrecognized     []byte `json:"-"`
}

func (m *TangramSessionEvent) Reset()                    { *m = TangramSessionEvent{} }
func (m *TangramSessionEvent) String() string            { return proto.CompactTextString(m) }
func (*TangramSessionEvent) ProtoMessage()               {}
func (*TangramSessionEvent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{0} }

func (m *TangramSessionEvent) GetLastDifficulty() int32 {
	if m != nil && m.LastDifficulty != nil {
		return *m.LastDifficulty
	}
	return 0
}

func (m *TangramSessionEvent) GetNumStartedTreasure() int32 {
	if m != nil && m.NumStartedTreasure != nil {
		return *m.NumStartedTreasure
	}
	return 0
}

func (m *TangramSessionEvent) GetNumSolvedTreasure() int32 {
	if m != nil && m.NumSolvedTreasure != nil {
		return *m.NumSolvedTreasure
	}
	return 0
}

func (m *TangramSessionEvent) GetNumAbandonedTreasure() int32 {
	if m != nil && m.NumAbandonedTreasure != nil {
		return *m.NumAbandonedTreasure
	}
	return 0
}

func (m *TangramSessionEvent) GetNumStartedClassic() int32 {
	if m != nil && m.NumStartedClassic != nil {
		return *m.NumStartedClassic
	}
	return 0
}

func (m *TangramSessionEvent) GetNumSolvedClassic() int32 {
	if m != nil && m.NumSolvedClassic != nil {
		return *m.NumSolvedClassic
	}
	return 0
}

func (m *TangramSessionEvent) GetNumAbandonedClassic() int32 {
	if m != nil && m.NumAbandonedClassic != nil {
		return *m.NumAbandonedClassic
	}
	return 0
}

func (m *TangramSessionEvent) GetNumStartedAcademy() int32 {
	if m != nil && m.NumStartedAcademy != nil {
		return *m.NumStartedAcademy
	}
	return 0
}

func (m *TangramSessionEvent) GetNumSolvedAcademy() int32 {
	if m != nil && m.NumSolvedAcademy != nil {
		return *m.NumSolvedAcademy
	}
	return 0
}

var E_TangramSessionEvent_Tangram = &proto.ExtensionDesc{
	ExtendedType:  (*SessionEvent)(nil),
	ExtensionType: (*TangramSessionEvent)(nil),
	Field:         1000,
	Name:          "generated.TangramSessionEvent.tangram",
	Tag:           "bytes,1000,opt,name=tangram",
}

// Tangram level start event
type TangramLevelEvent struct {
	// identifying information
	Category   *string `protobuf:"bytes,1,opt,name=category" json:"category,omitempty"`
	Shape      *string `protobuf:"bytes,2,opt,name=shape" json:"shape,omitempty"`
	Difficulty *int32  `protobuf:"varint,3,opt,name=difficulty" json:"difficulty,omitempty"`
	CastleId   *string `protobuf:"bytes,4,opt,name=castle_id,json=castleId" json:"castle_id,omitempty"`
	Mode       *string `protobuf:"bytes,5,opt,name=mode" json:"mode,omitempty"`
	IntroLevel *int32  `protobuf:"varint,6,opt,name=intro_level,json=introLevel" json:"intro_level,omitempty"`
	// stats about what happened in the level
	HintsUsed               *int32 `protobuf:"varint,7,opt,name=hints_used,json=hintsUsed" json:"hints_used,omitempty"`
	FovHintsShown           *int32 `protobuf:"varint,8,opt,name=fov_hints_shown,json=fovHintsShown" json:"fov_hints_shown,omitempty"`
	FovHintMs               *int32 `protobuf:"varint,9,opt,name=fov_hint_ms,json=fovHintMs" json:"fov_hint_ms,omitempty"`
	TableHintsShown         *int32 `protobuf:"varint,10,opt,name=table_hints_shown,json=tableHintsShown" json:"table_hints_shown,omitempty"`
	ParallelogramHintsShown *int32 `protobuf:"varint,11,opt,name=parallelogram_hints_shown,json=parallelogramHintsShown" json:"parallelogram_hints_shown,omitempty"`
	XXX_unrecognized        []byte `json:"-"`
}

func (m *TangramLevelEvent) Reset()                    { *m = TangramLevelEvent{} }
func (m *TangramLevelEvent) String() string            { return proto.CompactTextString(m) }
func (*TangramLevelEvent) ProtoMessage()               {}
func (*TangramLevelEvent) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{1} }

func (m *TangramLevelEvent) GetCategory() string {
	if m != nil && m.Category != nil {
		return *m.Category
	}
	return ""
}

func (m *TangramLevelEvent) GetShape() string {
	if m != nil && m.Shape != nil {
		return *m.Shape
	}
	return ""
}

func (m *TangramLevelEvent) GetDifficulty() int32 {
	if m != nil && m.Difficulty != nil {
		return *m.Difficulty
	}
	return 0
}

func (m *TangramLevelEvent) GetCastleId() string {
	if m != nil && m.CastleId != nil {
		return *m.CastleId
	}
	return ""
}

func (m *TangramLevelEvent) GetMode() string {
	if m != nil && m.Mode != nil {
		return *m.Mode
	}
	return ""
}

func (m *TangramLevelEvent) GetIntroLevel() int32 {
	if m != nil && m.IntroLevel != nil {
		return *m.IntroLevel
	}
	return 0
}

func (m *TangramLevelEvent) GetHintsUsed() int32 {
	if m != nil && m.HintsUsed != nil {
		return *m.HintsUsed
	}
	return 0
}

func (m *TangramLevelEvent) GetFovHintsShown() int32 {
	if m != nil && m.FovHintsShown != nil {
		return *m.FovHintsShown
	}
	return 0
}

func (m *TangramLevelEvent) GetFovHintMs() int32 {
	if m != nil && m.FovHintMs != nil {
		return *m.FovHintMs
	}
	return 0
}

func (m *TangramLevelEvent) GetTableHintsShown() int32 {
	if m != nil && m.TableHintsShown != nil {
		return *m.TableHintsShown
	}
	return 0
}

func (m *TangramLevelEvent) GetParallelogramHintsShown() int32 {
	if m != nil && m.ParallelogramHintsShown != nil {
		return *m.ParallelogramHintsShown
	}
	return 0
}

var E_TangramLevelEvent_Tangram = &proto.ExtensionDesc{
	ExtendedType:  (*LevelEvent)(nil),
	ExtensionType: (*TangramLevelEvent)(nil),
	Field:         1000,
	Name:          "generated.TangramLevelEvent.tangram",
	Tag:           "bytes,1000,opt,name=tangram",
}

// Tangram specific per device properties
type TangramUserProperties struct {
	LastDifficulty          *int32 `protobuf:"varint,1,opt,name=last_difficulty,json=lastDifficulty" json:"last_difficulty,omitempty"`
	HintsUsed               *int32 `protobuf:"varint,2,opt,name=hints_used,json=hintsUsed" json:"hints_used,omitempty"`
	FovHintsShown           *int32 `protobuf:"varint,3,opt,name=fov_hints_shown,json=fovHintsShown" json:"fov_hints_shown,omitempty"`
	FovHintMs               *int32 `protobuf:"varint,4,opt,name=fov_hint_ms,json=fovHintMs" json:"fov_hint_ms,omitempty"`
	TableHintsShown         *int32 `protobuf:"varint,5,opt,name=table_hints_shown,json=tableHintsShown" json:"table_hints_shown,omitempty"`
	ParallelogramHintsShown *int32 `protobuf:"varint,6,opt,name=parallelogram_hints_shown,json=parallelogramHintsShown" json:"parallelogram_hints_shown,omitempty"`
	XXX_unrecognized        []byte `json:"-"`
}

func (m *TangramUserProperties) Reset()                    { *m = TangramUserProperties{} }
func (m *TangramUserProperties) String() string            { return proto.CompactTextString(m) }
func (*TangramUserProperties) ProtoMessage()               {}
func (*TangramUserProperties) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{2} }

func (m *TangramUserProperties) GetLastDifficulty() int32 {
	if m != nil && m.LastDifficulty != nil {
		return *m.LastDifficulty
	}
	return 0
}

func (m *TangramUserProperties) GetHintsUsed() int32 {
	if m != nil && m.HintsUsed != nil {
		return *m.HintsUsed
	}
	return 0
}

func (m *TangramUserProperties) GetFovHintsShown() int32 {
	if m != nil && m.FovHintsShown != nil {
		return *m.FovHintsShown
	}
	return 0
}

func (m *TangramUserProperties) GetFovHintMs() int32 {
	if m != nil && m.FovHintMs != nil {
		return *m.FovHintMs
	}
	return 0
}

func (m *TangramUserProperties) GetTableHintsShown() int32 {
	if m != nil && m.TableHintsShown != nil {
		return *m.TableHintsShown
	}
	return 0
}

func (m *TangramUserProperties) GetParallelogramHintsShown() int32 {
	if m != nil && m.ParallelogramHintsShown != nil {
		return *m.ParallelogramHintsShown
	}
	return 0
}

var E_TangramUserProperties_Tangram = &proto.ExtensionDesc{
	ExtendedType:  (*UserProperties)(nil),
	ExtensionType: (*TangramUserProperties)(nil),
	Field:         1000,
	Name:          "generated.TangramUserProperties.tangram",
	Tag:           "bytes,1000,opt,name=tangram",
}

// Tangram player state
type TangramPlayerState struct {
	TreasureAchievementsCompleted *int32 `protobuf:"varint,1,opt,name=treasure_achievements_completed,json=treasureAchievementsCompleted" json:"treasure_achievements_completed,omitempty"`
	ClassicAchievementsCompleted  *int32 `protobuf:"varint,2,opt,name=classic_achievements_completed,json=classicAchievementsCompleted" json:"classic_achievements_completed,omitempty"`
	SharedAchievementsCompleted   *int32 `protobuf:"varint,3,opt,name=shared_achievements_completed,json=sharedAchievementsCompleted" json:"shared_achievements_completed,omitempty"`
	BeatIntro                     *bool  `protobuf:"varint,4,opt,name=beat_intro,json=beatIntro" json:"beat_intro,omitempty"`
	IntroLevel                    *int32 `protobuf:"varint,5,opt,name=intro_level,json=introLevel" json:"intro_level,omitempty"`
	CurrentHints                  *int32 `protobuf:"varint,6,opt,name=current_hints,json=currentHints" json:"current_hints,omitempty"`
	TreasureHints                 *int32 `protobuf:"varint,7,opt,name=treasure_hints,json=treasureHints" json:"treasure_hints,omitempty"`
	TreasureHintsTotal            *int32 `protobuf:"varint,8,opt,name=treasure_hints_total,json=treasureHintsTotal" json:"treasure_hints_total,omitempty"`
	// number of puzzles solved at each difficulty level in classic mode
	ClassicSolved0 *int32 `protobuf:"varint,9,opt,name=classic_solved0,json=classicSolved0" json:"classic_solved0,omitempty"`
	ClassicSolved1 *int32 `protobuf:"varint,10,opt,name=classic_solved1,json=classicSolved1" json:"classic_solved1,omitempty"`
	ClassicSolved2 *int32 `protobuf:"varint,11,opt,name=classic_solved2,json=classicSolved2" json:"classic_solved2,omitempty"`
	ClassicSolved3 *int32 `protobuf:"varint,12,opt,name=classic_solved3,json=classicSolved3" json:"classic_solved3,omitempty"`
	// number of puzzles solved at each difficulty level in the new treasure mode
	TreasureSolved0           *int32 `protobuf:"varint,13,opt,name=treasure_solved0,json=treasureSolved0" json:"treasure_solved0,omitempty"`
	TreasureSolved1           *int32 `protobuf:"varint,14,opt,name=treasure_solved1,json=treasureSolved1" json:"treasure_solved1,omitempty"`
	TreasureSolved2           *int32 `protobuf:"varint,15,opt,name=treasure_solved2,json=treasureSolved2" json:"treasure_solved2,omitempty"`
	TreasureSolved3           *int32 `protobuf:"varint,16,opt,name=treasure_solved3,json=treasureSolved3" json:"treasure_solved3,omitempty"`
	TreasureSolvedUnique      *int32 `protobuf:"varint,17,opt,name=treasure_solved_unique,json=treasureSolvedUnique" json:"treasure_solved_unique,omitempty"`
	CastlesCompletedUnique    *int32 `protobuf:"varint,18,opt,name=castles_completed_unique,json=castlesCompletedUnique" json:"castles_completed_unique,omitempty"`
	BigCastlesCompletedUnique *int32 `protobuf:"varint,19,opt,name=big_castles_completed_unique,json=bigCastlesCompletedUnique" json:"big_castles_completed_unique,omitempty"`
	TreasuresCollectedUnique  *int32 `protobuf:"varint,20,opt,name=treasures_collected_unique,json=treasuresCollectedUnique" json:"treasures_collected_unique,omitempty"`
	PeopleRescuedUnique       *int32 `protobuf:"varint,21,opt,name=people_rescued_unique,json=peopleRescuedUnique" json:"people_rescued_unique,omitempty"`
	HasEnteredCastle          *bool  `protobuf:"varint,22,opt,name=has_entered_castle,json=hasEnteredCastle" json:"has_entered_castle,omitempty"`
	RescuedHomeless           *int32 `protobuf:"varint,23,opt,name=rescued_homeless,json=rescuedHomeless" json:"rescued_homeless,omitempty"`
	ClassicEnabled            *bool  `protobuf:"varint,24,opt,name=classic_enabled,json=classicEnabled" json:"classic_enabled,omitempty"`
	XXX_unrecognized          []byte `json:"-"`
}

func (m *TangramPlayerState) Reset()                    { *m = TangramPlayerState{} }
func (m *TangramPlayerState) String() string            { return proto.CompactTextString(m) }
func (*TangramPlayerState) ProtoMessage()               {}
func (*TangramPlayerState) Descriptor() ([]byte, []int) { return fileDescriptor1, []int{3} }

func (m *TangramPlayerState) GetTreasureAchievementsCompleted() int32 {
	if m != nil && m.TreasureAchievementsCompleted != nil {
		return *m.TreasureAchievementsCompleted
	}
	return 0
}

func (m *TangramPlayerState) GetClassicAchievementsCompleted() int32 {
	if m != nil && m.ClassicAchievementsCompleted != nil {
		return *m.ClassicAchievementsCompleted
	}
	return 0
}

func (m *TangramPlayerState) GetSharedAchievementsCompleted() int32 {
	if m != nil && m.SharedAchievementsCompleted != nil {
		return *m.SharedAchievementsCompleted
	}
	return 0
}

func (m *TangramPlayerState) GetBeatIntro() bool {
	if m != nil && m.BeatIntro != nil {
		return *m.BeatIntro
	}
	return false
}

func (m *TangramPlayerState) GetIntroLevel() int32 {
	if m != nil && m.IntroLevel != nil {
		return *m.IntroLevel
	}
	return 0
}

func (m *TangramPlayerState) GetCurrentHints() int32 {
	if m != nil && m.CurrentHints != nil {
		return *m.CurrentHints
	}
	return 0
}

func (m *TangramPlayerState) GetTreasureHints() int32 {
	if m != nil && m.TreasureHints != nil {
		return *m.TreasureHints
	}
	return 0
}

func (m *TangramPlayerState) GetTreasureHintsTotal() int32 {
	if m != nil && m.TreasureHintsTotal != nil {
		return *m.TreasureHintsTotal
	}
	return 0
}

func (m *TangramPlayerState) GetClassicSolved0() int32 {
	if m != nil && m.ClassicSolved0 != nil {
		return *m.ClassicSolved0
	}
	return 0
}

func (m *TangramPlayerState) GetClassicSolved1() int32 {
	if m != nil && m.ClassicSolved1 != nil {
		return *m.ClassicSolved1
	}
	return 0
}

func (m *TangramPlayerState) GetClassicSolved2() int32 {
	if m != nil && m.ClassicSolved2 != nil {
		return *m.ClassicSolved2
	}
	return 0
}

func (m *TangramPlayerState) GetClassicSolved3() int32 {
	if m != nil && m.ClassicSolved3 != nil {
		return *m.ClassicSolved3
	}
	return 0
}

func (m *TangramPlayerState) GetTreasureSolved0() int32 {
	if m != nil && m.TreasureSolved0 != nil {
		return *m.TreasureSolved0
	}
	return 0
}

func (m *TangramPlayerState) GetTreasureSolved1() int32 {
	if m != nil && m.TreasureSolved1 != nil {
		return *m.TreasureSolved1
	}
	return 0
}

func (m *TangramPlayerState) GetTreasureSolved2() int32 {
	if m != nil && m.TreasureSolved2 != nil {
		return *m.TreasureSolved2
	}
	return 0
}

func (m *TangramPlayerState) GetTreasureSolved3() int32 {
	if m != nil && m.TreasureSolved3 != nil {
		return *m.TreasureSolved3
	}
	return 0
}

func (m *TangramPlayerState) GetTreasureSolvedUnique() int32 {
	if m != nil && m.TreasureSolvedUnique != nil {
		return *m.TreasureSolvedUnique
	}
	return 0
}

func (m *TangramPlayerState) GetCastlesCompletedUnique() int32 {
	if m != nil && m.CastlesCompletedUnique != nil {
		return *m.CastlesCompletedUnique
	}
	return 0
}

func (m *TangramPlayerState) GetBigCastlesCompletedUnique() int32 {
	if m != nil && m.BigCastlesCompletedUnique != nil {
		return *m.BigCastlesCompletedUnique
	}
	return 0
}

func (m *TangramPlayerState) GetTreasuresCollectedUnique() int32 {
	if m != nil && m.TreasuresCollectedUnique != nil {
		return *m.TreasuresCollectedUnique
	}
	return 0
}

func (m *TangramPlayerState) GetPeopleRescuedUnique() int32 {
	if m != nil && m.PeopleRescuedUnique != nil {
		return *m.PeopleRescuedUnique
	}
	return 0
}

func (m *TangramPlayerState) GetHasEnteredCastle() bool {
	if m != nil && m.HasEnteredCastle != nil {
		return *m.HasEnteredCastle
	}
	return false
}

func (m *TangramPlayerState) GetRescuedHomeless() int32 {
	if m != nil && m.RescuedHomeless != nil {
		return *m.RescuedHomeless
	}
	return 0
}

func (m *TangramPlayerState) GetClassicEnabled() bool {
	if m != nil && m.ClassicEnabled != nil {
		return *m.ClassicEnabled
	}
	return false
}

var E_TangramPlayerState_Tangram = &proto.ExtensionDesc{
	ExtendedType:  (*PlayerState)(nil),
	ExtensionType: (*TangramPlayerState)(nil),
	Field:         1000,
	Name:          "generated.TangramPlayerState.tangram",
	Tag:           "bytes,1000,opt,name=tangram",
}

func init() {
	proto.RegisterType((*TangramSessionEvent)(nil), "generated.TangramSessionEvent")
	proto.RegisterType((*TangramLevelEvent)(nil), "generated.TangramLevelEvent")
	proto.RegisterType((*TangramUserProperties)(nil), "generated.TangramUserProperties")
	proto.RegisterType((*TangramPlayerState)(nil), "generated.TangramPlayerState")
	proto.RegisterExtension(E_TangramSessionEvent_Tangram)
	proto.RegisterExtension(E_TangramLevelEvent_Tangram)
	proto.RegisterExtension(E_TangramUserProperties_Tangram)
	proto.RegisterExtension(E_TangramPlayerState_Tangram)
}

func init() { proto.RegisterFile("TangramAnalyticsEvent.proto", fileDescriptor1) }

var fileDescriptor1 = []byte{
	// 957 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0x9c, 0x96, 0xeb, 0x4e, 0xdc, 0x46,
	0x14, 0xc7, 0x15, 0x60, 0x13, 0x7c, 0x08, 0xb0, 0x0c, 0xb7, 0xe1, 0x9a, 0x28, 0x55, 0xaf, 0xaa,
	0x10, 0x98, 0x7e, 0xa8, 0xaa, 0x4a, 0x15, 0x10, 0xaa, 0x44, 0x6a, 0x2b, 0xb4, 0x90, 0x7e, 0xb5,
	0x06, 0x7b, 0x60, 0x2d, 0x79, 0x3d, 0xdb, 0x99, 0x31, 0x15, 0xaf, 0xd2, 0x07, 0xeb, 0x2b, 0xb4,
	0x52, 0x5f, 0xa2, 0x73, 0xb5, 0xc7, 0x6b, 0xa7, 0x42, 0xfd, 0xc6, 0x9e, 0xf3, 0xfb, 0x1f, 0x1f,
	0xce, 0xcd, 0x86, 0xbd, 0x1b, 0x52, 0xde, 0x73, 0x32, 0x39, 0x2b, 0x49, 0xf1, 0x28, 0xf3, 0x54,
	0x5c, 0x3e, 0xd0, 0x52, 0x1e, 0x4d, 0x39, 0x93, 0x0c, 0x45, 0xf7, 0xb4, 0xa4, 0x9c, 0x48, 0x9a,
	0xed, 0x6e, 0xf4, 0x01, 0x6f, 0xfe, 0x58, 0x80, 0x75, 0x17, 0xe0, 0x9a, 0x0a, 0x91, 0xb3, 0xd2,
	0x78, 0xd1, 0xe7, 0xb0, 0x5a, 0x10, 0x21, 0x93, 0x2c, 0xbf, 0xbb, 0xcb, 0xd3, 0xaa, 0x90, 0x8f,
	0xf8, 0xd9, 0xeb, 0x67, 0x5f, 0x0c, 0x46, 0x2b, 0xda, 0xfc, 0xb6, 0xb6, 0xa2, 0x63, 0xd8, 0x28,
	0xab, 0x49, 0x22, 0x24, 0xe1, 0xea, 0x29, 0x89, 0xe4, 0x94, 0x88, 0x8a, 0x53, 0x3c, 0x67, 0x68,
	0xa4, 0x7c, 0xd7, 0xd6, 0x75, 0xe3, 0x3c, 0xe8, 0x08, 0xd6, 0x8d, 0x82, 0x15, 0x0f, 0xa1, 0x60,
	0xde, 0x08, 0xd6, 0xb4, 0xc0, 0x78, 0x6a, 0xfe, 0x1b, 0xd8, 0xd2, 0x3c, 0xb9, 0x25, 0x65, 0xc6,
	0xca, 0x50, 0xb2, 0x60, 0x24, 0xfa, 0xf9, 0x67, 0xde, 0xd9, 0x79, 0x8a, 0xcb, 0x2b, 0x55, 0x69,
	0x8b, 0x3c, 0xc5, 0x83, 0xe6, 0x29, 0xd6, 0x73, 0x61, 0x1d, 0xe8, 0x6b, 0x40, 0x41, 0x56, 0x1e,
	0x7f, 0x6e, 0xf0, 0x61, 0x9d, 0x94, 0xa7, 0x63, 0xd8, 0x6c, 0xe7, 0xe4, 0x05, 0x2f, 0x8c, 0x60,
	0x3d, 0x4c, 0xc9, 0x6b, 0x66, 0x32, 0x22, 0x29, 0xc9, 0xe8, 0xe4, 0x11, 0x2f, 0xce, 0x66, 0x74,
	0x66, 0x1d, 0x33, 0x19, 0x79, 0x3c, 0x9a, 0xc9, 0xc8, 0xd1, 0xf1, 0x08, 0x5e, 0x48, 0xdb, 0x47,
	0xb4, 0x7d, 0x54, 0x77, 0xfd, 0x28, 0x6c, 0x2a, 0xfe, 0x5b, 0x67, 0xb7, 0x14, 0x1f, 0x06, 0xfe,
	0x9e, 0xde, 0x8f, 0x7c, 0xa0, 0x37, 0x7f, 0xce, 0xc3, 0x9a, 0x03, 0x7e, 0xa2, 0x0f, 0xb4, 0xb0,
	0xa3, 0xb1, 0x0b, 0x8b, 0xa9, 0xd2, 0xde, 0x33, 0x6e, 0x67, 0x22, 0x1a, 0xd5, 0xbf, 0xd1, 0x06,
	0x0c, 0xc4, 0x98, 0x4c, 0x6d, 0xfb, 0xa3, 0x91, 0xfd, 0x81, 0x0e, 0x01, 0x82, 0x39, 0xb2, 0x8d,
	0x0e, 0x2c, 0x68, 0x0f, 0xa2, 0x54, 0x4d, 0x55, 0x41, 0x93, 0x3c, 0x33, 0x4d, 0x35, 0x21, 0xb5,
	0xe1, 0x7d, 0x86, 0x10, 0x2c, 0x4c, 0x58, 0x46, 0x4d, 0xe7, 0xa2, 0x91, 0xf9, 0x1b, 0xbd, 0x82,
	0xa5, 0xbc, 0x94, 0x9c, 0x25, 0x85, 0x4e, 0xcb, 0x75, 0x09, 0x8c, 0xc9, 0x24, 0x8a, 0x0e, 0x00,
	0xc6, 0xea, 0xa7, 0x48, 0x2a, 0x41, 0x33, 0xd7, 0x94, 0xc8, 0x58, 0x3e, 0x28, 0x03, 0xfa, 0x0c,
	0x56, 0xef, 0xd8, 0x43, 0x62, 0x11, 0x31, 0x66, 0xbf, 0x97, 0xae, 0x0d, 0xcb, 0xca, 0xfc, 0x4e,
	0x5b, 0xaf, 0xb5, 0x51, 0x25, 0xbe, 0xe4, 0xb9, 0x64, 0x22, 0x5c, 0xed, 0x23, 0xc7, 0xfc, 0x2c,
	0xd0, 0x57, 0xb0, 0x26, 0xc9, 0xad, 0xca, 0x3b, 0x8c, 0x04, 0x86, 0x5a, 0x35, 0x8e, 0x20, 0xd6,
	0x77, 0xb0, 0x33, 0x25, 0x9c, 0x14, 0x05, 0x2d, 0x98, 0xae, 0x68, 0x4b, 0xb3, 0x64, 0x34, 0xdb,
	0x2d, 0xa0, 0xd1, 0xc6, 0xbf, 0x34, 0xcd, 0xdd, 0x0c, 0x9a, 0xd7, 0x34, 0xc5, 0xb7, 0x76, 0xbf,
	0xdb, 0xda, 0x06, 0x6a, 0x1a, 0xfb, 0xd7, 0x1c, 0x6c, 0x3a, 0xb7, 0xaa, 0x07, 0xbf, 0xe2, 0x6c,
	0x4a, 0xb9, 0xcc, 0xa9, 0x78, 0xfa, 0xde, 0xb7, 0x2b, 0x3c, 0xf7, 0x84, 0x0a, 0xcf, 0x3f, 0xa1,
	0xc2, 0x0b, 0x4f, 0xaa, 0xf0, 0xe0, 0x7f, 0x54, 0xf8, 0xf9, 0x7f, 0x57, 0xf8, 0xd7, 0xa6, 0xc2,
	0x3b, 0x41, 0x0d, 0xdb, 0xd5, 0xf1, 0x55, 0x7e, 0xdd, 0xad, 0x72, 0x1b, 0x6c, 0x2a, 0xfd, 0x4f,
	0x04, 0xc8, 0x21, 0x57, 0x05, 0x79, 0xa4, 0x5c, 0x2d, 0xb9, 0xa4, 0xe8, 0x47, 0x78, 0xe5, 0xaf,
	0x98, 0xda, 0xec, 0x71, 0xae, 0x7a, 0x34, 0xa1, 0x3a, 0xd9, 0x94, 0x4d, 0xa6, 0x05, 0x55, 0x81,
	0x5d, 0xd9, 0x0f, 0x3c, 0x76, 0x16, 0x50, 0x17, 0x1e, 0x42, 0x6f, 0xe1, 0xd0, 0x5d, 0x9e, 0x8f,
	0x85, 0xb1, 0x9d, 0xd9, 0x77, 0x54, 0x7f, 0x94, 0x73, 0x38, 0x50, 0x8b, 0xca, 0xcd, 0x95, 0xe9,
	0x0d, 0x62, 0x5b, 0xb7, 0x67, 0xa1, 0xfe, 0x18, 0x6a, 0x1e, 0x6e, 0x29, 0x91, 0x89, 0x59, 0x42,
	0xd3, 0xc7, 0xc5, 0x51, 0xa4, 0x2d, 0xef, 0xb5, 0x61, 0x76, 0x63, 0x07, 0x9d, 0x8d, 0xfd, 0x04,
	0x96, 0xd3, 0x8a, 0x73, 0x15, 0xd4, 0xb6, 0xcd, 0x35, 0xec, 0xa5, 0x33, 0x9a, 0x56, 0xa1, 0x4f,
	0x61, 0xa5, 0x2e, 0x9b, 0xa5, 0xec, 0x6a, 0x2f, 0x7b, 0xab, 0xc5, 0xd4, 0x3b, 0xa9, 0x8d, 0x25,
	0x92, 0x49, 0x52, 0xb8, 0x1d, 0x47, 0x2d, 0xf8, 0x46, 0x7b, 0xf4, 0xd8, 0xfb, 0x3a, 0xda, 0x7b,
	0x7b, 0xec, 0x96, 0x7d, 0xc5, 0x99, 0xed, 0xb1, 0x3d, 0xee, 0x82, 0x27, 0x6e, 0xdf, 0xdb, 0xe0,
	0x49, 0x17, 0x8c, 0xdd, 0x92, 0xb7, 0xc1, 0xb8, 0x0b, 0x9e, 0xe2, 0x97, 0x3d, 0xe0, 0x29, 0xfa,
	0x12, 0x86, 0xf5, 0x7f, 0xe5, 0x93, 0x5c, 0x76, 0x9b, 0xe0, 0xec, 0x3e, 0xcb, 0x2e, 0x7a, 0x82,
	0x57, 0xfa, 0xd0, 0x93, 0x1e, 0x34, 0xc6, 0xab, 0x7d, 0x68, 0xdc, 0x83, 0x9e, 0xe2, 0x61, 0x1f,
	0x7a, 0xaa, 0xdf, 0xd9, 0x33, 0x68, 0x52, 0x95, 0xf9, 0x6f, 0x15, 0xc5, 0x6b, 0xf6, 0x9d, 0xdd,
	0x16, 0x7c, 0x30, 0x3e, 0xf4, 0x2d, 0x60, 0x7b, 0xf6, 0x83, 0xd9, 0xf3, 0x3a, 0x64, 0x74, 0x5b,
	0xce, 0x5f, 0xcf, 0x9d, 0x53, 0xfe, 0x00, 0xfb, 0xb7, 0xf9, 0x7d, 0xf2, 0x51, 0xf5, 0xba, 0x51,
	0xef, 0x28, 0xe6, 0xa2, 0x3f, 0xc0, 0xf7, 0xb0, 0xeb, 0x53, 0xd2, 0x72, 0x75, 0x22, 0xd2, 0x40,
	0xbe, 0x61, 0xe4, 0xb8, 0x26, 0x2e, 0x3c, 0xe0, 0xd4, 0xea, 0x73, 0x60, 0x4a, 0x99, 0x0a, 0x98,
	0x28, 0x67, 0x5a, 0x35, 0xc2, 0x4d, 0xfb, 0x39, 0x60, 0x9d, 0x23, 0xeb, 0x73, 0x1a, 0xf5, 0x7a,
	0x1f, 0x13, 0x91, 0xa8, 0xd9, 0xa6, 0x7a, 0xf3, 0x6c, 0xea, 0x78, 0xcb, 0x2c, 0xce, 0x50, 0x79,
	0x2e, 0xad, 0xc3, 0xe6, 0xab, 0x6b, 0xef, 0x43, 0x8f, 0xd9, 0x84, 0xaa, 0xff, 0x40, 0xe0, 0x6d,
	0x5b, 0x7b, 0x67, 0x7f, 0xe7, 0xcc, 0xe1, 0x40, 0xd1, 0x52, 0x9f, 0xc8, 0x0c, 0x63, 0x13, 0xd5,
	0x0f, 0xd4, 0xa5, 0xb5, 0xc6, 0x57, 0xcd, 0xcd, 0xdb, 0x0a, 0x2e, 0x5a, 0x70, 0xa7, 0xfc, 0xc1,
	0x3b, 0xe8, 0x1e, 0xbc, 0x80, 0xaa, 0xaf, 0xdd, 0x39, 0x3a, 0x1f, 0xd6, 0x5f, 0x99, 0x8e, 0xfb,
	0x37, 0x00, 0x00, 0xff, 0xff, 0x01, 0x63, 0x05, 0x3a, 0xa0, 0x0a, 0x00, 0x00,
}
