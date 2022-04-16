package xivapi

import (
	"context"
	"fmt"
	"github.com/google/go-cmp/cmp"
	"net/http"
	"testing"
)

func TestServersService_Servers(t *testing.T) {
	client, mux, _, teardown := setup()
	defer teardown()

	mux.HandleFunc("/servers", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `["Adamantoise","Aegis","Alexander","Anima","Asura","Atomos","Bahamut","Balmung","Behemoth","Belias","Brynhildr","Cactuar","Carbuncle","Cerberus","Chocobo","Coeurl","Diabolos","Durandal","Excalibur","Exodus","Faerie","Famfrit","Fenrir","Garuda","Gilgamesh","Goblin","Gungnir","Hades","Hyperion","Ifrit","Ixion","Jenova","Kujata","Lamia","Leviathan","Lich","Louisoix","Malboro","Mandragora","Masamune","Mateus","Midgardsormr","Moogle","Odin","Omega","Pandaemonium","Phoenix","Ragnarok","Ramuh","Ridill","Sargatanas","Shinryu","Shiva","Siren","Tiamat","Titan","Tonberry","Typhon","Ultima","Ultros","Unicorn","Valefor","Yojimbo","Zalera","Zeromus","Zodiark","Spriggan","Twintania","Bismarck","Ravana","Sephirot","Sophia","Zurvan","HongYuHai","ShenYiZhiDi","LaNuoXiYa","HuanYingQunDao","MengYaChi","YuZhouHeYin","WoXianXiRan","ChenXiWangZuo","BaiYinXiang","BaiJinHuanXiang","ShenQuanHen","ChaoFengTing","LvRenZhanQiao","FuXiaoZhiJian","Longchaoshendian","MengYuBaoJing","ZiShuiZhanQiao","YanXia","JingYuZhuangYuan","MoDuNa","HaiMaoChaWu","RouFengHaiWan","HuPoYuan","ShuiJingTa2","YinLeiHu2","TaiYangHaiAn2","YiXiuJiaDe2","HongChaChuan2"]`)
	})

	ctx := context.Background()
	result, _, err := client.Servers.Servers(ctx)
	if err != nil {
		t.Errorf("Search.Items returned error: %v", err)
	}
	var want = &ServersResponse{"Adamantoise", "Aegis", "Alexander", "Anima", "Asura", "Atomos", "Bahamut", "Balmung", "Behemoth", "Belias", "Brynhildr", "Cactuar", "Carbuncle", "Cerberus", "Chocobo", "Coeurl", "Diabolos", "Durandal", "Excalibur", "Exodus", "Faerie", "Famfrit", "Fenrir", "Garuda", "Gilgamesh", "Goblin", "Gungnir", "Hades", "Hyperion", "Ifrit", "Ixion", "Jenova", "Kujata", "Lamia", "Leviathan", "Lich", "Louisoix", "Malboro", "Mandragora", "Masamune", "Mateus", "Midgardsormr", "Moogle", "Odin", "Omega", "Pandaemonium", "Phoenix", "Ragnarok", "Ramuh", "Ridill", "Sargatanas", "Shinryu", "Shiva", "Siren", "Tiamat", "Titan", "Tonberry", "Typhon", "Ultima", "Ultros", "Unicorn", "Valefor", "Yojimbo", "Zalera", "Zeromus", "Zodiark", "Spriggan", "Twintania", "Bismarck", "Ravana", "Sephirot", "Sophia", "Zurvan", "HongYuHai", "ShenYiZhiDi", "LaNuoXiYa", "HuanYingQunDao", "MengYaChi", "YuZhouHeYin", "WoXianXiRan", "ChenXiWangZuo", "BaiYinXiang", "BaiJinHuanXiang", "ShenQuanHen", "ChaoFengTing", "LvRenZhanQiao", "FuXiaoZhiJian", "Longchaoshendian", "MengYuBaoJing", "ZiShuiZhanQiao", "YanXia", "JingYuZhuangYuan", "MoDuNa", "HaiMaoChaWu", "RouFengHaiWan", "HuPoYuan", "ShuiJingTa2", "YinLeiHu2", "TaiYangHaiAn2", "YiXiuJiaDe2", "HongChaChuan2"}

	if !cmp.Equal(result, want) {
		t.Errorf("Servers returned %+v, want %+v", result, want)
	}
}
