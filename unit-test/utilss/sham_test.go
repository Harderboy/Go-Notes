package utils

import "testing"

func TestFunc1(t *testing.T) {
    t.Log("TestFunc1 run")
}
func TestFunc2(t *testing.T) {
    info := "TestFunc12 run"
    t.Log(info + "===test log the method name `Log`")
    t.Logf("output : %s", info+"===test log the method name `Logf`")
    t.Error(info + "===test log the method name `Error`")
    t.Errorf("output : %s", info+"===test log the method name `Errorf`")
    t.Fatal(info + "===test log the method name `Fatal`")
    t.Fatalf("output : %s", info+"===test log the method name `Fatalf`")
}