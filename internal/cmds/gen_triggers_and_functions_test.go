// Code generated DO NOT EDIT

package cmds

import "testing"

func triggers_and_functions0(s Builder) {
	s.Tfcall().LibraryFunction("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Build()
	s.Tfcall().LibraryFunction("1").Numkeys(1).Key("1").Key("1").Build()
	s.Tfcall().LibraryFunction("1").Numkeys(1).Arg("1").Arg("1").Build()
	s.Tfcall().LibraryFunction("1").Numkeys(1).Build()
	s.Tfcallasync().LibraryFunction("1").Numkeys(1).Key("1").Key("1").Arg("1").Arg("1").Build()
	s.Tfcallasync().LibraryFunction("1").Numkeys(1).Key("1").Key("1").Build()
	s.Tfcallasync().LibraryFunction("1").Numkeys(1).Arg("1").Arg("1").Build()
	s.Tfcallasync().LibraryFunction("1").Numkeys(1).Build()
	s.TfunctionDelete().LibraryName("1").Build()
	s.TfunctionList().LibraryName("1").Withcode().Verbose().V().Build()
	s.TfunctionList().LibraryName("1").Withcode().Verbose().Build()
	s.TfunctionList().LibraryName("1").Withcode().V().Build()
	s.TfunctionList().LibraryName("1").Withcode().Build()
	s.TfunctionList().LibraryName("1").Verbose().Build()
	s.TfunctionList().LibraryName("1").V().Build()
	s.TfunctionList().LibraryName("1").Build()
	s.TfunctionList().Withcode().Build()
	s.TfunctionList().Verbose().Build()
	s.TfunctionList().V().Build()
	s.TfunctionList().Build()
	s.TfunctionLoad().Replace().Config("1").LibraryCode("1").Build()
	s.TfunctionLoad().Replace().LibraryCode("1").Build()
	s.TfunctionLoad().Config("1").LibraryCode("1").Build()
	s.TfunctionLoad().LibraryCode("1").Build()
}

func TestCommand_InitSlot_triggers_and_functions(t *testing.T) {
	var s = NewBuilder(InitSlot)
	t.Run("0", func(t *testing.T) { triggers_and_functions0(s) })
}

func TestCommand_NoSlot_triggers_and_functions(t *testing.T) {
	var s = NewBuilder(NoSlot)
	t.Run("0", func(t *testing.T) { triggers_and_functions0(s) })
}
