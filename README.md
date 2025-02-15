## 设计思路

```go
package custom_plugin


type BotInfos struct{}

func (b *BotInfos) GetScope() uint32 {
	return define.GroupScope
}

var _ plugin_tree.PluginInterface = &BotInfos{}

func (b *BotInfos) GetPluginInfo() string {
	return "BotInfos -> bot相关信息展示\n${func | fn | f | 功能清单 | functions | 功能}"
}
func (b *BotInfos) GetPaths() []string {

	prefix := "@ " + define.BotQQ + " "
	return []string{
		prefix + "f",
		prefix + "func",
		prefix + "fn",
		prefix + "功能清单",
		prefix + "functions",
		prefix + "功能",
	}
}
func (b *BotInfos) GetPluginHandler() plugin_tree.PluginHandler {
	return func(api *bot_action.BotActionAPI, ctx *models.MessageContext) plugin_tree.ContextResult {
		resStr := "功能清单：\n"
		for i, plugin := range plugin_tree.CustomPlugins {
			resStr += fmt.Sprintf("[%d]: %s\n\n", i+1, plugin.GetPluginInfo())
		}

		groupId := ctx.MessageChain.GetTargetId()
		api.SendGroupMessage(
			models.NewGroupChain(groupId).Text(resStr),
			func(messageId int64) {
				global_data.BotMessageIdStack.GetStack(groupId).Push(messageId)
			})
		return plugin_tree.ContextResult{}
	}
}

```



```
[/Script/Pal.PalGameWorldSettings]
OptionSettings=(Difficulty=None,RandomizerType=None,RandomizerSeed="",DayTimeSpeedRate=1,NightTimeSpeedRate=1,ExpRate=20,PalCaptureRate=2,PalSpawnNumRate=1.5,PalDamageRateAttack=1,PalDamageRateDefense=1,PlayerDamageRateAttack=1,PlayerDamageRateDefense=0.5,PlayerStomachDecreaceRate=0.1,PlayerStaminaDecreaceRate=0.1,PlayerAutoHPRegeneRate=2,PlayerAutoHpRegeneRateInSleep=2,PalStomachDecreaceRate=0.1,PalStaminaDecreaceRate=0.5,PalAutoHPRegeneRate=1,PalAutoHpRegeneRateInSleep=1,BuildObjectHpRate=1,BuildObjectDamageRate=1,BuildObjectDeteriorationDamageRate=1,CollectionDropRate=1,CollectionObjectHpRate=1,CollectionObjectRespawnSpeedRate=1,EnemyDropItemRate=3,DeathPenalty=None,bEnablePlayerToPlayerDamage=False,bEnableFriendlyFire=False,bEnableInvaderEnemy=True,bActiveUNKO=False,bEnableAimAssistPad=True,bEnableAimAssistKeyboard=True,DropItemMaxNum=3000,DropItemMaxNum_UNKO=100,BaseCampMaxNum=128,BaseCampWorkerMaxNum=20,DropItemAliveMaxHours=1,bAutoResetGuildNoOnlinePlayers=False,AutoResetGuildTimeNoOnlinePlayers=300,GuildPlayerMaxNum=20,BaseCampMaxNumInGuild=20,PalEggDefaultHatchingTime=0,WorkSpeedRate=1,AutoSaveSpan=30,bIsMultiplay=True,bIsPvP=False,bHardcore=False,bPalLost=False,bCanPickupOtherGuildDeathPenaltyDrop=False,bEnableNonLoginPenalty=False,bEnableFastTravel=True,bIsStartLocationSelectByMap=True,bExistPlayerAfterLogout=False,bEnableDefenseOtherGuildPlayer=False,bInvisibleOtherGuildBaseCampAreaFX=False,bBuildAreaLimit=False,ItemWeightRate=0,CoopPlayerMaxNum=4,ServerPlayerMaxNum=32,ServerName="Default Palworld Server",ServerDescription="",AdminPassword="",ServerPassword="",PublicPort=8211,PublicIP="",RCONEnabled=False,RCONPort=25575,Region="",bUseAuth=True,BanListURL="https://api.palworldgame.com/api/banlist.txt",RESTAPIEnabled=False,RESTAPIPort=8212,bShowPlayerList=False,ChatPostLimitPerMinute=10,AllowConnectPlatform=Steam,bIsUseBackupSaveData=True,LogFormatType=Text,SupplyDropSpan=90,EnablePredatorBossPal=True,MaxBuildingLimitNum="0",ServerReplicatePawnCullDistance=15000)

```

