package monster

import "testing"

// 测试用例
func TestStore(t *testing.T) {
	monster := &Monster{
		Name:  "小红",
		Age:   10,
		Skill: "222",
	}
	res := monster.Store()
	if !res {
		t.Fatalf("monster.Store()错误，希望是%v 实际是%v", true, res)
	}
	t.Logf("monster.Store() 测试成功")
}
func TestReStore(t *testing.T) {
	var monster Monster
	res := monster.ReStore()
	if !res {
		t.Fatalf("monster.ReStore()错误，希望是%v 实际是%v", true, res)
	}
	if monster.Name != "小红" {
		t.Fatalf("monster.ReStore()错误，希望是%v 实际是%v", "小红", res)
	}
	t.Logf("monster.ReStore() 测试成功")
}
