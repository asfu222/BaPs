package proto

type Protocol int32

const (
	Protocol_Common_Cheat                              Protocol = -9999
	Protocol_Error                                     Protocol = -1
	Protocol_None                                      Protocol = 0
	Protocol_System_Version                            Protocol = 1
	Protocol_Session_Info                              Protocol = 2
	Protocol_NetworkTime_Sync                          Protocol = 3
	Protocol_NetworkTime_SyncReply                     Protocol = 4
	Protocol_Audit_GachaStatistics                     Protocol = 5
	Protocol_Account_Create                            Protocol = 1000
	Protocol_Account_Nickname                          Protocol = 1001
	Protocol_Account_Auth                              Protocol = 1002
	Protocol_Account_CurrencySync                      Protocol = 1003
	Protocol_Account_SetRepresentCharacterAndComment   Protocol = 1004
	Protocol_Account_GetTutorial                       Protocol = 1005
	Protocol_Account_SetTutorial                       Protocol = 1006
	Protocol_Account_PassCheck                         Protocol = 1007
	Protocol_Account_VerifyForYostar                   Protocol = 1008
	Protocol_Account_CheckYostar                       Protocol = 1009
	Protocol_Account_CallName                          Protocol = 1010
	Protocol_Account_BirthDay                          Protocol = 1011
	Protocol_Account_Auth2                             Protocol = 1012
	Protocol_Account_LinkReward                        Protocol = 1013
	Protocol_Account_ReportXignCodeCheater             Protocol = 1014
	Protocol_Account_DismissRepurchasablePopup         Protocol = 1015
	Protocol_Account_InvalidateToken                   Protocol = 1016
	Protocol_Account_LoginSync                         Protocol = 1017
	Protocol_Account_Reset                             Protocol = 1018
	Protocol_Account_RequestBirthdayMail               Protocol = 1019
	Protocol_Account_CheckAccountLevelReward           Protocol = 1020
	Protocol_Account_ReceiveAccountLevelReward         Protocol = 1021
	Protocol_Character_List                            Protocol = 2000
	Protocol_Character_Transcendence                   Protocol = 2001
	Protocol_Character_ExpGrowth                       Protocol = 2002
	Protocol_Character_FavorGrowth                     Protocol = 2003
	Protocol_Character_UpdateSkillLevel                Protocol = 2004
	Protocol_Character_UnlockWeapon                    Protocol = 2005
	Protocol_Character_WeaponExpGrowth                 Protocol = 2006
	Protocol_Character_WeaponTranscendence             Protocol = 2007
	Protocol_Character_SetFavorites                    Protocol = 2008
	Protocol_Character_SetCostume                      Protocol = 2009
	Protocol_Character_BatchSkillLevelUpdate           Protocol = 2010
	Protocol_Character_PotentialGrowth                 Protocol = 2011
	Protocol_Equipment_List                            Protocol = 3000
	Protocol_Equipment_Sell                            Protocol = 3001
	Protocol_Equipment_Equip                           Protocol = 3002
	Protocol_Equipment_LevelUp                         Protocol = 3003
	Protocol_Equipment_TierUp                          Protocol = 3004
	Protocol_Equipment_Lock                            Protocol = 3005
	Protocol_Equipment_BatchGrowth                     Protocol = 3006
	Protocol_Item_List                                 Protocol = 4000
	Protocol_Item_Sell                                 Protocol = 4001
	Protocol_Item_Consume                              Protocol = 4002
	Protocol_Item_Lock                                 Protocol = 4003
	Protocol_Item_BulkConsume                          Protocol = 4004
	Protocol_Item_SelectTicket                         Protocol = 4005
	Protocol_Item_AutoSynth                            Protocol = 4006
	Protocol_Echelon_List                              Protocol = 5000
	Protocol_Echelon_Save                              Protocol = 5001
	Protocol_Echelon_PresetList                        Protocol = 5002
	Protocol_Echelon_PresetSave                        Protocol = 5003
	Protocol_Echelon_PresetGroupRename                 Protocol = 5004
	Protocol_Campaign_List                             Protocol = 6000
	Protocol_Campaign_EnterMainStage                   Protocol = 6001
	Protocol_Campaign_ConfirmMainStage                 Protocol = 6002
	Protocol_Campaign_DeployEchelon                    Protocol = 6003
	Protocol_Campaign_WithdrawEchelon                  Protocol = 6004
	Protocol_Campaign_MapMove                          Protocol = 6005
	Protocol_Campaign_EndTurn                          Protocol = 6006
	Protocol_Campaign_EnterTactic                      Protocol = 6007
	Protocol_Campaign_TacticResult                     Protocol = 6008
	Protocol_Campaign_Retreat                          Protocol = 6009
	Protocol_Campaign_ChapterClearReward               Protocol = 6010
	Protocol_Campaign_Heal                             Protocol = 6011
	Protocol_Campaign_EnterSubStage                    Protocol = 6012
	Protocol_Campaign_SubStageResult                   Protocol = 6013
	Protocol_Campaign_Portal                           Protocol = 6014
	Protocol_Campaign_ConfirmTutorialStage             Protocol = 6015
	Protocol_Campaign_PurchasePlayCountHardStage       Protocol = 6016
	Protocol_Campaign_EnterTutorialStage               Protocol = 6017
	Protocol_Campaign_TutorialStageResult              Protocol = 6018
	Protocol_Campaign_RestartMainStage                 Protocol = 6019
	Protocol_Campaign_EnterMainStageStrategySkip       Protocol = 6020
	Protocol_Campaign_MainStageStrategySkipResult      Protocol = 6021
	Protocol_Mail_List                                 Protocol = 7000
	Protocol_Mail_Check                                Protocol = 7001
	Protocol_Mail_Receive                              Protocol = 7002
	Protocol_Mission_List                              Protocol = 8000
	Protocol_Mission_Reward                            Protocol = 8001
	Protocol_Mission_MultipleReward                    Protocol = 8002
	Protocol_Mission_GuideReward                       Protocol = 8003
	Protocol_Mission_MultipleGuideReward               Protocol = 8004
	Protocol_Mission_Sync                              Protocol = 8005
	Protocol_Mission_GuideMissionSeasonList            Protocol = 8006
	Protocol_Attendance_List                           Protocol = 9000
	Protocol_Attendance_Check                          Protocol = 9001
	Protocol_Attendance_Reward                         Protocol = 9002
	Protocol_Shop_BuyMerchandise                       Protocol = 10000
	Protocol_Shop_BuyGacha                             Protocol = 10001
	Protocol_Shop_List                                 Protocol = 10002
	Protocol_Shop_Refresh                              Protocol = 10003
	Protocol_Shop_BuyEligma                            Protocol = 10004
	Protocol_Shop_BuyGacha2                            Protocol = 10005
	Protocol_Shop_GachaRecruitList                     Protocol = 10006
	Protocol_Shop_BuyRefreshMerchandise                Protocol = 10007
	Protocol_Shop_BuyGacha3                            Protocol = 10008
	Protocol_Shop_BuyAP                                Protocol = 10009
	Protocol_Shop_BeforehandGachaGet                   Protocol = 10010
	Protocol_Shop_BeforehandGachaRun                   Protocol = 10011
	Protocol_Shop_BeforehandGachaSave                  Protocol = 10012
	Protocol_Shop_BeforehandGachaPick                  Protocol = 10013
	Protocol_Shop_PickupSelectionGachaGet              Protocol = 10014
	Protocol_Shop_PickupSelectionGachaSet              Protocol = 10015
	Protocol_Shop_PickupSelectionGachaBuy              Protocol = 10016
	Protocol_Recipe_Craft                              Protocol = 11000
	Protocol_MemoryLobby_List                          Protocol = 12000
	Protocol_MemoryLobby_SetMain                       Protocol = 12001
	Protocol_MemoryLobby_UpdateLobbyMode               Protocol = 12002
	Protocol_MemoryLobby_Interact                      Protocol = 12003
	Protocol_CumulativeTimeReward_List                 Protocol = 13000
	Protocol_CumulativeTimeReward_Reward               Protocol = 13001
	Protocol_OpenCondition_List                        Protocol = 15000
	Protocol_OpenCondition_Set                         Protocol = 15001
	Protocol_OpenCondition_EventList                   Protocol = 15002
	Protocol_Toast_List                                Protocol = 16000
	Protocol_Raid_List                                 Protocol = 17000
	Protocol_Raid_CompleteList                         Protocol = 17001
	Protocol_Raid_Detail                               Protocol = 17002
	Protocol_Raid_Search                               Protocol = 17003
	Protocol_Raid_CreateBattle                         Protocol = 17004
	Protocol_Raid_EnterBattle                          Protocol = 17005
	Protocol_Raid_BattleUpdate                         Protocol = 17006
	Protocol_Raid_EndBattle                            Protocol = 17007
	Protocol_Raid_Reward                               Protocol = 17008
	Protocol_Raid_RewardAll                            Protocol = 17009
	Protocol_Raid_Revive                               Protocol = 17010
	Protocol_Raid_Share                                Protocol = 17011
	Protocol_Raid_SeasonInfo                           Protocol = 17012
	Protocol_Raid_SeasonReward                         Protocol = 17013
	Protocol_Raid_Lobby                                Protocol = 17014
	Protocol_Raid_GiveUp                               Protocol = 17015
	Protocol_Raid_OpponentList                         Protocol = 17016
	Protocol_Raid_RankingReward                        Protocol = 17017
	Protocol_Raid_Login                                Protocol = 17018
	Protocol_Raid_Sweep                                Protocol = 17019
	Protocol_Raid_GetBestTeam                          Protocol = 17020
	Protocol_Raid_RankingIndex                         Protocol = 17021
	Protocol_SkipHistory_List                          Protocol = 18000
	Protocol_SkipHistory_Save                          Protocol = 18001
	Protocol_Scenario_List                             Protocol = 19000
	Protocol_Scenario_Clear                            Protocol = 19001
	Protocol_Scenario_GroupHistoryUpdate               Protocol = 19002
	Protocol_Scenario_Skip                             Protocol = 19003
	Protocol_Scenario_Select                           Protocol = 19004
	Protocol_Scenario_AccountStudentChange             Protocol = 19005
	Protocol_Scenario_LobbyStudentChange               Protocol = 19006
	Protocol_Scenario_SpecialLobbyChange               Protocol = 19007
	Protocol_Scenario_Enter                            Protocol = 19008
	Protocol_Scenario_EnterMainStage                   Protocol = 19009
	Protocol_Scenario_ConfirmMainStage                 Protocol = 19010
	Protocol_Scenario_DeployEchelon                    Protocol = 19011
	Protocol_Scenario_WithdrawEchelon                  Protocol = 19012
	Protocol_Scenario_MapMove                          Protocol = 19013
	Protocol_Scenario_EndTurn                          Protocol = 19014
	Protocol_Scenario_EnterTactic                      Protocol = 19015
	Protocol_Scenario_TacticResult                     Protocol = 19016
	Protocol_Scenario_Retreat                          Protocol = 19017
	Protocol_Scenario_Portal                           Protocol = 19018
	Protocol_Scenario_RestartMainStage                 Protocol = 19019
	Protocol_Scenario_SkipMainStage                    Protocol = 19020
	Protocol_Cafe_Get                                  Protocol = 20000
	Protocol_Cafe_Ack                                  Protocol = 20001
	Protocol_Cafe_Deploy                               Protocol = 20002
	Protocol_Cafe_Relocate                             Protocol = 20003
	Protocol_Cafe_Remove                               Protocol = 20004
	Protocol_Cafe_RemoveAll                            Protocol = 20005
	Protocol_Cafe_Interact                             Protocol = 20006
	Protocol_Cafe_ListPreset                           Protocol = 20007
	Protocol_Cafe_RenamePreset                         Protocol = 20008
	Protocol_Cafe_ClearPreset                          Protocol = 20009
	Protocol_Cafe_UpdatePresetFurniture                Protocol = 20010
	Protocol_Cafe_ApplyPreset                          Protocol = 20011
	Protocol_Cafe_RankUp                               Protocol = 20012
	Protocol_Cafe_ReceiveCurrency                      Protocol = 20013
	Protocol_Cafe_GiveGift                             Protocol = 20014
	Protocol_Cafe_SummonCharacter                      Protocol = 20015
	Protocol_Cafe_TrophyHistory                        Protocol = 20016
	Protocol_Cafe_ApplyTemplate                        Protocol = 20017
	Protocol_Cafe_Open                                 Protocol = 20018
	Protocol_Cafe_Travel                               Protocol = 20019
	Protocol_Craft_List                                Protocol = 21000
	Protocol_Craft_SelectNode                          Protocol = 21001
	Protocol_Craft_UpdateNodeLevel                     Protocol = 21002
	Protocol_Craft_BeginProcess                        Protocol = 21003
	Protocol_Craft_CompleteProcess                     Protocol = 21004
	Protocol_Craft_Reward                              Protocol = 21005
	Protocol_Craft_HistoryList                         Protocol = 21006
	Protocol_Craft_ShiftingBeginProcess                Protocol = 21007
	Protocol_Craft_ShiftingCompleteProcess             Protocol = 21008
	Protocol_Craft_ShiftingReward                      Protocol = 21009
	Protocol_Craft_AutoBeginProcess                    Protocol = 21010
	Protocol_Craft_CompleteProcessAll                  Protocol = 21011
	Protocol_Craft_RewardAll                           Protocol = 21012
	Protocol_Craft_ShiftingCompleteProcessAll          Protocol = 21013
	Protocol_Craft_ShiftingRewardAll                   Protocol = 21014
	Protocol_Arena_EnterLobby                          Protocol = 22000
	Protocol_Arena_Login                               Protocol = 22001
	Protocol_Arena_SettingChange                       Protocol = 22002
	Protocol_Arena_OpponentList                        Protocol = 22003
	Protocol_Arena_EnterBattle                         Protocol = 22004
	Protocol_Arena_EnterBattlePart1                    Protocol = 22005
	Protocol_Arena_EnterBattlePart2                    Protocol = 22006
	Protocol_Arena_BattleResult                        Protocol = 22007
	Protocol_Arena_CumulativeTimeReward                Protocol = 22008
	Protocol_Arena_DailyReward                         Protocol = 22009
	Protocol_Arena_RankList                            Protocol = 22010
	Protocol_Arena_History                             Protocol = 22011
	Protocol_Arena_RecordSync                          Protocol = 22012
	Protocol_Arena_TicketPurchase                      Protocol = 22013
	Protocol_Arena_DamageReport                        Protocol = 22014
	Protocol_Arena_CheckSeasonCloseReward              Protocol = 22015
	Protocol_Arena_SyncEchelonSettingTime              Protocol = 22016
	Protocol_WeekDungeon_List                          Protocol = 23000
	Protocol_WeekDungeon_EnterBattle                   Protocol = 23001
	Protocol_WeekDungeon_BattleResult                  Protocol = 23002
	Protocol_WeekDungeon_Retreat                       Protocol = 23003
	Protocol_Academy_GetInfo                           Protocol = 24000
	Protocol_Academy_AttendSchedule                    Protocol = 24001
	Protocol_Academy_AttendFavorSchedule               Protocol = 24002
	Protocol_Event_GetList                             Protocol = 25000
	Protocol_Event_GetImage                            Protocol = 25001
	Protocol_Event_UseCoupon                           Protocol = 25002
	Protocol_Event_RewardIncrease                      Protocol = 25003
	Protocol_ContentSave_Get                           Protocol = 26000
	Protocol_ContentSave_Discard                       Protocol = 26001
	Protocol_ContentSweep_Request                      Protocol = 27000
	Protocol_ContentSweep_MultiSweep                   Protocol = 27001
	Protocol_ContentSweep_MultiSweepPresetList         Protocol = 27002
	Protocol_ContentSweep_SetMultiSweepPreset          Protocol = 27003
	Protocol_ContentSweep_SetMultiSweepPresetName      Protocol = 27004
	Protocol_Clan_Lobby                                Protocol = 28000
	Protocol_Clan_Login                                Protocol = 28001
	Protocol_Clan_Search                               Protocol = 28002
	Protocol_Clan_Create                               Protocol = 28003
	Protocol_Clan_Member                               Protocol = 28004
	Protocol_Clan_Applicant                            Protocol = 28005
	Protocol_Clan_Join                                 Protocol = 28006
	Protocol_Clan_Quit                                 Protocol = 28007
	Protocol_Clan_Permit                               Protocol = 28008
	Protocol_Clan_Kick                                 Protocol = 28009
	Protocol_Clan_Setting                              Protocol = 28010
	Protocol_Clan_Confer                               Protocol = 28011
	Protocol_Clan_Dismiss                              Protocol = 28012
	Protocol_Clan_AutoJoin                             Protocol = 28013
	Protocol_Clan_MemberList                           Protocol = 28014
	Protocol_Clan_CancelApply                          Protocol = 28015
	Protocol_Clan_MyAssistList                         Protocol = 28016
	Protocol_Clan_SetAssist                            Protocol = 28017
	Protocol_Clan_ChatLog                              Protocol = 28018
	Protocol_Clan_Check                                Protocol = 28019
	Protocol_Clan_AllAssistList                        Protocol = 28020
	Protocol_Billing_TransactionStartByYostar          Protocol = 29000
	Protocol_Billing_TransactionEndByYostar            Protocol = 29001
	Protocol_Billing_PurchaseListByYostar              Protocol = 29002
	Protocol_Billing_PurchaseFreeProduct               Protocol = 29003
	Protocol_EventContent_AdventureList                Protocol = 30000
	Protocol_EventContent_EnterMainStage               Protocol = 30001
	Protocol_EventContent_ConfirmMainStage             Protocol = 30002
	Protocol_EventContent_EnterTactic                  Protocol = 30003
	Protocol_EventContent_TacticResult                 Protocol = 30004
	Protocol_EventContent_EnterSubStage                Protocol = 30005
	Protocol_EventContent_SubStageResult               Protocol = 30006
	Protocol_EventContent_DeployEchelon                Protocol = 30007
	Protocol_EventContent_WithdrawEchelon              Protocol = 30008
	Protocol_EventContent_MapMove                      Protocol = 30009
	Protocol_EventContent_EndTurn                      Protocol = 30010
	Protocol_EventContent_Retreat                      Protocol = 30011
	Protocol_EventContent_Portal                       Protocol = 30012
	Protocol_EventContent_PurchasePlayCountHardStage   Protocol = 30013
	Protocol_EventContent_ShopList                     Protocol = 30014
	Protocol_EventContent_ShopRefresh                  Protocol = 30015
	Protocol_EventContent_ReceiveStageTotalReward      Protocol = 30016
	Protocol_EventContent_EnterMainGroundStage         Protocol = 30017
	Protocol_EventContent_MainGroundStageResult        Protocol = 30018
	Protocol_EventContent_ShopBuyMerchandise           Protocol = 30019
	Protocol_EventContent_ShopBuyRefreshMerchandise    Protocol = 30020
	Protocol_EventContent_SelectBuff                   Protocol = 30021
	Protocol_EventContent_BoxGachaShopList             Protocol = 30022
	Protocol_EventContent_BoxGachaShopPurchase         Protocol = 30023
	Protocol_EventContent_BoxGachaShopRefresh          Protocol = 30024
	Protocol_EventContent_CollectionList               Protocol = 30025
	Protocol_EventContent_CollectionForMission         Protocol = 30026
	Protocol_EventContent_ScenarioGroupHistoryUpdate   Protocol = 30027
	Protocol_EventContent_CardShopList                 Protocol = 30028
	Protocol_EventContent_CardShopShuffle              Protocol = 30029
	Protocol_EventContent_CardShopPurchase             Protocol = 30030
	Protocol_EventContent_RestartMainStage             Protocol = 30031
	Protocol_EventContent_LocationGetInfo              Protocol = 30032
	Protocol_EventContent_LocationAttendSchedule       Protocol = 30033
	Protocol_EventContent_FortuneGachaPurchase         Protocol = 30034
	Protocol_EventContent_SubEventLobby                Protocol = 30035
	Protocol_EventContent_EnterStoryStage              Protocol = 30036
	Protocol_EventContent_StoryStageResult             Protocol = 30037
	Protocol_EventContent_DiceRaceLobby                Protocol = 30038
	Protocol_EventContent_DiceRaceRoll                 Protocol = 30039
	Protocol_EventContent_DiceRaceLapReward            Protocol = 30040
	Protocol_EventContent_PermanentList                Protocol = 30041
	Protocol_EventContent_DiceRaceUseItem              Protocol = 30042
	Protocol_EventContent_CardShopPurchaseAll          Protocol = 30043
	Protocol_EventContent_TreasureLobby                Protocol = 30044
	Protocol_EventContent_TreasureFlip                 Protocol = 30045
	Protocol_EventContent_TreasureNextRound            Protocol = 30046
	Protocol_TTS_GetFile                               Protocol = 31000
	Protocol_ContentLog_UIOpenStatistics               Protocol = 32000
	Protocol_MomoTalk_OutLine                          Protocol = 33000
	Protocol_MomoTalk_MessageList                      Protocol = 33001
	Protocol_MomoTalk_Read                             Protocol = 33002
	Protocol_MomoTalk_Reply                            Protocol = 33003
	Protocol_MomoTalk_FavorSchedule                    Protocol = 33004
	Protocol_ClearDeck_List                            Protocol = 34000
	Protocol_MiniGame_StageList                        Protocol = 35000
	Protocol_MiniGame_EnterStage                       Protocol = 35001
	Protocol_MiniGame_Result                           Protocol = 35002
	Protocol_MiniGame_MissionList                      Protocol = 35003
	Protocol_MiniGame_MissionReward                    Protocol = 35004
	Protocol_MiniGame_MissionMultipleReward            Protocol = 35005
	Protocol_MiniGame_ShootingLobby                    Protocol = 35006
	Protocol_MiniGame_ShootingBattleEnter              Protocol = 35007
	Protocol_MiniGame_ShootingBattleResult             Protocol = 35008
	Protocol_MiniGame_ShootingSweep                    Protocol = 35009
	Protocol_MiniGame_TableBoardSync                   Protocol = 35010
	Protocol_MiniGame_TableBoardMove                   Protocol = 35011
	Protocol_MiniGame_TableBoardEncounterInput         Protocol = 35012
	Protocol_MiniGame_TableBoardBattleEncounter        Protocol = 35013
	Protocol_MiniGame_TableBoardBattleRunAway          Protocol = 35014
	Protocol_MiniGame_TableBoardClearThema             Protocol = 35015
	Protocol_MiniGame_TableBoardUseItem                Protocol = 35016
	Protocol_MiniGame_TableBoardResurrect              Protocol = 35017
	Protocol_MiniGame_TableBoardSweep                  Protocol = 35018
	Protocol_MiniGame_TableBoardMoveThema              Protocol = 35019
	Protocol_MiniGame_DreamMakerGetInfo                Protocol = 35020
	Protocol_MiniGame_DreamMakerNewGame                Protocol = 35021
	Protocol_MiniGame_DreamMakerRestart                Protocol = 35022
	Protocol_MiniGame_DreamMakerAttendSchedule         Protocol = 35023
	Protocol_MiniGame_DreamMakerDailyClosing           Protocol = 35024
	Protocol_MiniGame_DreamMakerEnding                 Protocol = 35025
	Protocol_MiniGame_DefenseGetInfo                   Protocol = 35026
	Protocol_MiniGame_DefenseEnterBattle               Protocol = 35027
	Protocol_MiniGame_DefenseBattleResult              Protocol = 35028
	Protocol_MiniGame_RoadPuzzleGetInfo                Protocol = 35029
	Protocol_MiniGame_RoadPuzzleTilePlace              Protocol = 35030
	Protocol_MiniGame_RoadPuzzleSaveStage              Protocol = 35031
	Protocol_MiniGame_RoadPuzzleClearStage             Protocol = 35032
	Protocol_MiniGame_CCGLobby                         Protocol = 35033
	Protocol_MiniGame_CCGCreateGame                    Protocol = 35034
	Protocol_MiniGame_CCGSweep                         Protocol = 35035
	Protocol_MiniGame_CCGEnterStage                    Protocol = 35036
	Protocol_MiniGame_CCGEndStageDual                  Protocol = 35037
	Protocol_MiniGame_CCGEndStageEvent                 Protocol = 35038
	Protocol_MiniGame_CCGSelectRewardCard              Protocol = 35039
	Protocol_Minigame_CCGReplaceCharacter              Protocol = 35040
	Protocol_MiniGame_CCGSelectCampAction              Protocol = 35041
	Protocol_MiniGame_CCGCompleteGame                  Protocol = 35042
	Protocol_MiniGame_CCGGiveupGame                    Protocol = 35043
	Protocol_MiniGame_CCGRerollReward                  Protocol = 35044
	Protocol_MiniGame_CCGBuyPerk                       Protocol = 35045
	Protocol_Notification_LobbyCheck                   Protocol = 36000
	Protocol_Notification_EventContentReddotCheck      Protocol = 36001
	Protocol_ProofToken_RequestQuestion                Protocol = 37000
	Protocol_ProofToken_Submit                         Protocol = 37001
	Protocol_SchoolDungeon_List                        Protocol = 38000
	Protocol_SchoolDungeon_EnterBattle                 Protocol = 38001
	Protocol_SchoolDungeon_BattleResult                Protocol = 38002
	Protocol_SchoolDungeon_Retreat                     Protocol = 38003
	Protocol_TimeAttackDungeon_Lobby                   Protocol = 39000
	Protocol_TimeAttackDungeon_CreateBattle            Protocol = 39001
	Protocol_TimeAttackDungeon_EnterBattle             Protocol = 39002
	Protocol_TimeAttackDungeon_EndBattle               Protocol = 39003
	Protocol_TimeAttackDungeon_Sweep                   Protocol = 39004
	Protocol_TimeAttackDungeon_GiveUp                  Protocol = 39005
	Protocol_TimeAttackDungeon_Login                   Protocol = 39006
	Protocol_WorldRaid_Lobby                           Protocol = 40000
	Protocol_WorldRaid_BossList                        Protocol = 40001
	Protocol_WorldRaid_EnterBattle                     Protocol = 40002
	Protocol_WorldRaid_BattleResult                    Protocol = 40003
	Protocol_WorldRaid_ReceiveReward                   Protocol = 40004
	Protocol_ResetableContent_Get                      Protocol = 41000
	Protocol_Conquest_GetInfo                          Protocol = 42000
	Protocol_Conquest_Conquer                          Protocol = 42001
	Protocol_Conquest_ConquerWithBattleStart           Protocol = 42002
	Protocol_Conquest_ConquerWithBattleResult          Protocol = 42003
	Protocol_Conquest_DeployEchelon                    Protocol = 42004
	Protocol_Conquest_ManageBase                       Protocol = 42005
	Protocol_Conquest_UpgradeBase                      Protocol = 42006
	Protocol_Conquest_TakeEventObject                  Protocol = 42007
	Protocol_Conquest_EventObjectBattleStart           Protocol = 42008
	Protocol_Conquest_EventObjectBattleResult          Protocol = 42009
	Protocol_Conquest_ReceiveCalculateRewards          Protocol = 42010
	Protocol_Conquest_NormalizeEchelon                 Protocol = 42011
	Protocol_Conquest_Check                            Protocol = 42012
	Protocol_Conquest_ErosionBattleStart               Protocol = 42013
	Protocol_Conquest_ErosionBattleResult              Protocol = 42014
	Protocol_Conquest_MainStoryGetInfo                 Protocol = 42015
	Protocol_Conquest_MainStoryConquer                 Protocol = 42016
	Protocol_Conquest_MainStoryConquerWithBattleStart  Protocol = 42017
	Protocol_Conquest_MainStoryConquerWithBattleResult Protocol = 42018
	Protocol_Conquest_MainStoryCheck                   Protocol = 42019
	Protocol_Friend_List                               Protocol = 43000
	Protocol_Friend_Remove                             Protocol = 43001
	Protocol_Friend_GetFriendDetailedInfo              Protocol = 43002
	Protocol_Friend_GetIdCard                          Protocol = 43003
	Protocol_Friend_SetIdCard                          Protocol = 43004
	Protocol_Friend_Search                             Protocol = 43005
	Protocol_Friend_SendFriendRequest                  Protocol = 43006
	Protocol_Friend_AcceptFriendRequest                Protocol = 43007
	Protocol_Friend_DeclineFriendRequest               Protocol = 43008
	Protocol_Friend_CancelFriendRequest                Protocol = 43009
	Protocol_Friend_Check                              Protocol = 43010
	Protocol_Friend_ListByIds                          Protocol = 43011
	Protocol_Friend_Block                              Protocol = 43012
	Protocol_Friend_Unblock                            Protocol = 43013
	Protocol_CharacterGear_List                        Protocol = 44000
	Protocol_CharacterGear_Unlock                      Protocol = 44001
	Protocol_CharacterGear_TierUp                      Protocol = 44002
	Protocol_EliminateRaid_Login                       Protocol = 45000
	Protocol_EliminateRaid_Lobby                       Protocol = 45001
	Protocol_EliminateRaid_OpponentList                Protocol = 45002
	Protocol_EliminateRaid_GetBestTeam                 Protocol = 45003
	Protocol_EliminateRaid_CreateBattle                Protocol = 45004
	Protocol_EliminateRaid_EnterBattle                 Protocol = 45005
	Protocol_EliminateRaid_EndBattle                   Protocol = 45006
	Protocol_EliminateRaid_GiveUp                      Protocol = 45007
	Protocol_EliminateRaid_Sweep                       Protocol = 45008
	Protocol_EliminateRaid_SeasonReward                Protocol = 45009
	Protocol_EliminateRaid_RankingReward               Protocol = 45010
	Protocol_EliminateRaid_LimitedReward               Protocol = 45011
	Protocol_EliminateRaid_RankingIndex                Protocol = 45012
	Protocol_Attachment_Get                            Protocol = 46000
	Protocol_Attachment_EmblemList                     Protocol = 46001
	Protocol_Attachment_EmblemAcquire                  Protocol = 46002
	Protocol_Attachment_EmblemAttach                   Protocol = 46003
	Protocol_Sticker_Login                             Protocol = 47000
	Protocol_Sticker_Lobby                             Protocol = 47001
	Protocol_Sticker_UseSticker                        Protocol = 47002
	Protocol_Field_Sync                                Protocol = 48000
	Protocol_Field_Interaction                         Protocol = 48001
	Protocol_Field_QuestClear                          Protocol = 48002
	Protocol_Field_SceneChanged                        Protocol = 48003
	Protocol_Field_EndDate                             Protocol = 48004
	Protocol_Field_EnterStage                          Protocol = 48005
	Protocol_Field_StageResult                         Protocol = 48006
	Protocol_MultiFloorRaid_Sync                       Protocol = 49000
	Protocol_MultiFloorRaid_EnterBattle                Protocol = 49001
	Protocol_MultiFloorRaid_EndBattle                  Protocol = 49002
	Protocol_MultiFloorRaid_ReceiveReward              Protocol = 49003
	Protocol_Queuing_GetTicket                         Protocol = 50000
	Protocol_BattlePass_GetInfo                        Protocol = 51000
	Protocol_BattlePass_BuyLevel                       Protocol = 51001
	Protocol_BattlePass_ReceiveReward                  Protocol = 51002
	Protocol_BattlePass_MissionList                    Protocol = 51003
	Protocol_BattlePass_MissionSingleReward            Protocol = 51004
	Protocol_BattlePass_MissionMultipleReward          Protocol = 51005
	Protocol_BattlePass_Check                          Protocol = 51006
	Protocol_GmTalk                                    Protocol = 90000
)

var (
	Protocol_name = map[int32]string{
		-9999: "Common_Cheat",
		-1:    "Error",
		0:     "None",
		1:     "System_Version",
		2:     "Session_Info",
		3:     "NetworkTime_Sync",
		4:     "NetworkTime_SyncReply",
		5:     "Audit_GachaStatistics",
		1000:  "Account_Create",
		1001:  "Account_Nickname",
		1002:  "Account_Auth",
		1003:  "Account_CurrencySync",
		1004:  "Account_SetRepresentCharacterAndComment",
		1005:  "Account_GetTutorial",
		1006:  "Account_SetTutorial",
		1007:  "Account_PassCheck",
		1008:  "Account_VerifyForYostar",
		1009:  "Account_CheckYostar",
		1010:  "Account_CallName",
		1011:  "Account_BirthDay",
		1012:  "Account_Auth2",
		1013:  "Account_LinkReward",
		1014:  "Account_ReportXignCodeCheater",
		1015:  "Account_DismissRepurchasablePopup",
		1016:  "Account_InvalidateToken",
		1017:  "Account_LoginSync",
		1018:  "Account_Reset",
		1019:  "Account_RequestBirthdayMail",
		1020:  "Account_CheckAccountLevelReward",
		1021:  "Account_ReceiveAccountLevelReward",
		2000:  "Character_List",
		2001:  "Character_Transcendence",
		2002:  "Character_ExpGrowth",
		2003:  "Character_FavorGrowth",
		2004:  "Character_UpdateSkillLevel",
		2005:  "Character_UnlockWeapon",
		2006:  "Character_WeaponExpGrowth",
		2007:  "Character_WeaponTranscendence",
		2008:  "Character_SetFavorites",
		2009:  "Character_SetCostume",
		2010:  "Character_BatchSkillLevelUpdate",
		2011:  "Character_PotentialGrowth",
		3000:  "Equipment_List",
		3001:  "Equipment_Sell",
		3002:  "Equipment_Equip",
		3003:  "Equipment_LevelUp",
		3004:  "Equipment_TierUp",
		3005:  "Equipment_Lock",
		3006:  "Equipment_BatchGrowth",
		4000:  "Item_List",
		4001:  "Item_Sell",
		4002:  "Item_Consume",
		4003:  "Item_Lock",
		4004:  "Item_BulkConsume",
		4005:  "Item_SelectTicket",
		4006:  "Item_AutoSynth",
		5000:  "Echelon_List",
		5001:  "Echelon_Save",
		5002:  "Echelon_PresetList",
		5003:  "Echelon_PresetSave",
		5004:  "Echelon_PresetGroupRename",
		6000:  "Campaign_List",
		6001:  "Campaign_EnterMainStage",
		6002:  "Campaign_ConfirmMainStage",
		6003:  "Campaign_DeployEchelon",
		6004:  "Campaign_WithdrawEchelon",
		6005:  "Campaign_MapMove",
		6006:  "Campaign_EndTurn",
		6007:  "Campaign_EnterTactic",
		6008:  "Campaign_TacticResult",
		6009:  "Campaign_Retreat",
		6010:  "Campaign_ChapterClearReward",
		6011:  "Campaign_Heal",
		6012:  "Campaign_EnterSubStage",
		6013:  "Campaign_SubStageResult",
		6014:  "Campaign_Portal",
		6015:  "Campaign_ConfirmTutorialStage",
		6016:  "Campaign_PurchasePlayCountHardStage",
		6017:  "Campaign_EnterTutorialStage",
		6018:  "Campaign_TutorialStageResult",
		6019:  "Campaign_RestartMainStage",
		6020:  "Campaign_EnterMainStageStrategySkip",
		6021:  "Campaign_MainStageStrategySkipResult",
		7000:  "Mail_List",
		7001:  "Mail_Check",
		7002:  "Mail_Receive",
		8000:  "Mission_List",
		8001:  "Mission_Reward",
		8002:  "Mission_MultipleReward",
		8003:  "Mission_GuideReward",
		8004:  "Mission_MultipleGuideReward",
		8005:  "Mission_Sync",
		8006:  "Mission_GuideMissionSeasonList",
		9000:  "Attendance_List",
		9001:  "Attendance_Check",
		9002:  "Attendance_Reward",
		10000: "Shop_BuyMerchandise",
		10001: "Shop_BuyGacha",
		10002: "Shop_List",
		10003: "Shop_Refresh",
		10004: "Shop_BuyEligma",
		10005: "Shop_BuyGacha2",
		10006: "Shop_GachaRecruitList",
		10007: "Shop_BuyRefreshMerchandise",
		10008: "Shop_BuyGacha3",
		10009: "Shop_BuyAP",
		10010: "Shop_BeforehandGachaGet",
		10011: "Shop_BeforehandGachaRun",
		10012: "Shop_BeforehandGachaSave",
		10013: "Shop_BeforehandGachaPick",
		10014: "Shop_PickupSelectionGachaGet",
		10015: "Shop_PickupSelectionGachaSet",
		10016: "Shop_PickupSelectionGachaBuy",
		11000: "Recipe_Craft",
		12000: "MemoryLobby_List",
		12001: "MemoryLobby_SetMain",
		12002: "MemoryLobby_UpdateLobbyMode",
		12003: "MemoryLobby_Interact",
		13000: "CumulativeTimeReward_List",
		13001: "CumulativeTimeReward_Reward",
		15000: "OpenCondition_List",
		15001: "OpenCondition_Set",
		15002: "OpenCondition_EventList",
		16000: "Toast_List",
		17000: "Raid_List",
		17001: "Raid_CompleteList",
		17002: "Raid_Detail",
		17003: "Raid_Search",
		17004: "Raid_CreateBattle",
		17005: "Raid_EnterBattle",
		17006: "Raid_BattleUpdate",
		17007: "Raid_EndBattle",
		17008: "Raid_Reward",
		17009: "Raid_RewardAll",
		17010: "Raid_Revive",
		17011: "Raid_Share",
		17012: "Raid_SeasonInfo",
		17013: "Raid_SeasonReward",
		17014: "Raid_Lobby",
		17015: "Raid_GiveUp",
		17016: "Raid_OpponentList",
		17017: "Raid_RankingReward",
		17018: "Raid_Login",
		17019: "Raid_Sweep",
		17020: "Raid_GetBestTeam",
		17021: "Raid_RankingIndex",
		18000: "SkipHistory_List",
		18001: "SkipHistory_Save",
		19000: "Scenario_List",
		19001: "Scenario_Clear",
		19002: "Scenario_GroupHistoryUpdate",
		19003: "Scenario_Skip",
		19004: "Scenario_Select",
		19005: "Scenario_AccountStudentChange",
		19006: "Scenario_LobbyStudentChange",
		19007: "Scenario_SpecialLobbyChange",
		19008: "Scenario_Enter",
		19009: "Scenario_EnterMainStage",
		19010: "Scenario_ConfirmMainStage",
		19011: "Scenario_DeployEchelon",
		19012: "Scenario_WithdrawEchelon",
		19013: "Scenario_MapMove",
		19014: "Scenario_EndTurn",
		19015: "Scenario_EnterTactic",
		19016: "Scenario_TacticResult",
		19017: "Scenario_Retreat",
		19018: "Scenario_Portal",
		19019: "Scenario_RestartMainStage",
		19020: "Scenario_SkipMainStage",
		20000: "Cafe_Get",
		20001: "Cafe_Ack",
		20002: "Cafe_Deploy",
		20003: "Cafe_Relocate",
		20004: "Cafe_Remove",
		20005: "Cafe_RemoveAll",
		20006: "Cafe_Interact",
		20007: "Cafe_ListPreset",
		20008: "Cafe_RenamePreset",
		20009: "Cafe_ClearPreset",
		20010: "Cafe_UpdatePresetFurniture",
		20011: "Cafe_ApplyPreset",
		20012: "Cafe_RankUp",
		20013: "Cafe_ReceiveCurrency",
		20014: "Cafe_GiveGift",
		20015: "Cafe_SummonCharacter",
		20016: "Cafe_TrophyHistory",
		20017: "Cafe_ApplyTemplate",
		20018: "Cafe_Open",
		20019: "Cafe_Travel",
		21000: "Craft_List",
		21001: "Craft_SelectNode",
		21002: "Craft_UpdateNodeLevel",
		21003: "Craft_BeginProcess",
		21004: "Craft_CompleteProcess",
		21005: "Craft_Reward",
		21006: "Craft_HistoryList",
		21007: "Craft_ShiftingBeginProcess",
		21008: "Craft_ShiftingCompleteProcess",
		21009: "Craft_ShiftingReward",
		21010: "Craft_AutoBeginProcess",
		21011: "Craft_CompleteProcessAll",
		21012: "Craft_RewardAll",
		21013: "Craft_ShiftingCompleteProcessAll",
		21014: "Craft_ShiftingRewardAll",
		22000: "Arena_EnterLobby",
		22001: "Arena_Login",
		22002: "Arena_SettingChange",
		22003: "Arena_OpponentList",
		22004: "Arena_EnterBattle",
		22005: "Arena_EnterBattlePart1",
		22006: "Arena_EnterBattlePart2",
		22007: "Arena_BattleResult",
		22008: "Arena_CumulativeTimeReward",
		22009: "Arena_DailyReward",
		22010: "Arena_RankList",
		22011: "Arena_History",
		22012: "Arena_RecordSync",
		22013: "Arena_TicketPurchase",
		22014: "Arena_DamageReport",
		22015: "Arena_CheckSeasonCloseReward",
		22016: "Arena_SyncEchelonSettingTime",
		23000: "WeekDungeon_List",
		23001: "WeekDungeon_EnterBattle",
		23002: "WeekDungeon_BattleResult",
		23003: "WeekDungeon_Retreat",
		24000: "Academy_GetInfo",
		24001: "Academy_AttendSchedule",
		24002: "Academy_AttendFavorSchedule",
		25000: "Event_GetList",
		25001: "Event_GetImage",
		25002: "Event_UseCoupon",
		25003: "Event_RewardIncrease",
		26000: "ContentSave_Get",
		26001: "ContentSave_Discard",
		27000: "ContentSweep_Request",
		27001: "ContentSweep_MultiSweep",
		27002: "ContentSweep_MultiSweepPresetList",
		27003: "ContentSweep_SetMultiSweepPreset",
		27004: "ContentSweep_SetMultiSweepPresetName",
		28000: "Clan_Lobby",
		28001: "Clan_Login",
		28002: "Clan_Search",
		28003: "Clan_Create",
		28004: "Clan_Member",
		28005: "Clan_Applicant",
		28006: "Clan_Join",
		28007: "Clan_Quit",
		28008: "Clan_Permit",
		28009: "Clan_Kick",
		28010: "Clan_Setting",
		28011: "Clan_Confer",
		28012: "Clan_Dismiss",
		28013: "Clan_AutoJoin",
		28014: "Clan_MemberList",
		28015: "Clan_CancelApply",
		28016: "Clan_MyAssistList",
		28017: "Clan_SetAssist",
		28018: "Clan_ChatLog",
		28019: "Clan_Check",
		28020: "Clan_AllAssistList",
		29000: "Billing_TransactionStartByYostar",
		29001: "Billing_TransactionEndByYostar",
		29002: "Billing_PurchaseListByYostar",
		29003: "Billing_PurchaseFreeProduct",
		30000: "EventContent_AdventureList",
		30001: "EventContent_EnterMainStage",
		30002: "EventContent_ConfirmMainStage",
		30003: "EventContent_EnterTactic",
		30004: "EventContent_TacticResult",
		30005: "EventContent_EnterSubStage",
		30006: "EventContent_SubStageResult",
		30007: "EventContent_DeployEchelon",
		30008: "EventContent_WithdrawEchelon",
		30009: "EventContent_MapMove",
		30010: "EventContent_EndTurn",
		30011: "EventContent_Retreat",
		30012: "EventContent_Portal",
		30013: "EventContent_PurchasePlayCountHardStage",
		30014: "EventContent_ShopList",
		30015: "EventContent_ShopRefresh",
		30016: "EventContent_ReceiveStageTotalReward",
		30017: "EventContent_EnterMainGroundStage",
		30018: "EventContent_MainGroundStageResult",
		30019: "EventContent_ShopBuyMerchandise",
		30020: "EventContent_ShopBuyRefreshMerchandise",
		30021: "EventContent_SelectBuff",
		30022: "EventContent_BoxGachaShopList",
		30023: "EventContent_BoxGachaShopPurchase",
		30024: "EventContent_BoxGachaShopRefresh",
		30025: "EventContent_CollectionList",
		30026: "EventContent_CollectionForMission",
		30027: "EventContent_ScenarioGroupHistoryUpdate",
		30028: "EventContent_CardShopList",
		30029: "EventContent_CardShopShuffle",
		30030: "EventContent_CardShopPurchase",
		30031: "EventContent_RestartMainStage",
		30032: "EventContent_LocationGetInfo",
		30033: "EventContent_LocationAttendSchedule",
		30034: "EventContent_FortuneGachaPurchase",
		30035: "EventContent_SubEventLobby",
		30036: "EventContent_EnterStoryStage",
		30037: "EventContent_StoryStageResult",
		30038: "EventContent_DiceRaceLobby",
		30039: "EventContent_DiceRaceRoll",
		30040: "EventContent_DiceRaceLapReward",
		30041: "EventContent_PermanentList",
		30042: "EventContent_DiceRaceUseItem",
		30043: "EventContent_CardShopPurchaseAll",
		30044: "EventContent_TreasureLobby",
		30045: "EventContent_TreasureFlip",
		30046: "EventContent_TreasureNextRound",
		31000: "TTS_GetFile",
		32000: "ContentLog_UIOpenStatistics",
		33000: "MomoTalk_OutLine",
		33001: "MomoTalk_MessageList",
		33002: "MomoTalk_Read",
		33003: "MomoTalk_Reply",
		33004: "MomoTalk_FavorSchedule",
		34000: "ClearDeck_List",
		35000: "MiniGame_StageList",
		35001: "MiniGame_EnterStage",
		35002: "MiniGame_Result",
		35003: "MiniGame_MissionList",
		35004: "MiniGame_MissionReward",
		35005: "MiniGame_MissionMultipleReward",
		35006: "MiniGame_ShootingLobby",
		35007: "MiniGame_ShootingBattleEnter",
		35008: "MiniGame_ShootingBattleResult",
		35009: "MiniGame_ShootingSweep",
		35010: "MiniGame_TableBoardSync",
		35011: "MiniGame_TableBoardMove",
		35012: "MiniGame_TableBoardEncounterInput",
		35013: "MiniGame_TableBoardBattleEncounter",
		35014: "MiniGame_TableBoardBattleRunAway",
		35015: "MiniGame_TableBoardClearThema",
		35016: "MiniGame_TableBoardUseItem",
		35017: "MiniGame_TableBoardResurrect",
		35018: "MiniGame_TableBoardSweep",
		35019: "MiniGame_TableBoardMoveThema",
		35020: "MiniGame_DreamMakerGetInfo",
		35021: "MiniGame_DreamMakerNewGame",
		35022: "MiniGame_DreamMakerRestart",
		35023: "MiniGame_DreamMakerAttendSchedule",
		35024: "MiniGame_DreamMakerDailyClosing",
		35025: "MiniGame_DreamMakerEnding",
		35026: "MiniGame_DefenseGetInfo",
		35027: "MiniGame_DefenseEnterBattle",
		35028: "MiniGame_DefenseBattleResult",
		35029: "MiniGame_RoadPuzzleGetInfo",
		35030: "MiniGame_RoadPuzzleTilePlace",
		35031: "MiniGame_RoadPuzzleSaveStage",
		35032: "MiniGame_RoadPuzzleClearStage",
		35033: "MiniGame_CCGLobby",
		35034: "MiniGame_CCGCreateGame",
		35035: "MiniGame_CCGSweep",
		35036: "MiniGame_CCGEnterStage",
		35037: "MiniGame_CCGEndStageDual",
		35038: "MiniGame_CCGEndStageEvent",
		35039: "MiniGame_CCGSelectRewardCard",
		35040: "Minigame_CCGReplaceCharacter",
		35041: "MiniGame_CCGSelectCampAction",
		35042: "MiniGame_CCGCompleteGame",
		35043: "MiniGame_CCGGiveupGame",
		35044: "MiniGame_CCGRerollReward",
		35045: "MiniGame_CCGBuyPerk",
		36000: "Notification_LobbyCheck",
		36001: "Notification_EventContentReddotCheck",
		37000: "ProofToken_RequestQuestion",
		37001: "ProofToken_Submit",
		38000: "SchoolDungeon_List",
		38001: "SchoolDungeon_EnterBattle",
		38002: "SchoolDungeon_BattleResult",
		38003: "SchoolDungeon_Retreat",
		39000: "TimeAttackDungeon_Lobby",
		39001: "TimeAttackDungeon_CreateBattle",
		39002: "TimeAttackDungeon_EnterBattle",
		39003: "TimeAttackDungeon_EndBattle",
		39004: "TimeAttackDungeon_Sweep",
		39005: "TimeAttackDungeon_GiveUp",
		39006: "TimeAttackDungeon_Login",
		40000: "WorldRaid_Lobby",
		40001: "WorldRaid_BossList",
		40002: "WorldRaid_EnterBattle",
		40003: "WorldRaid_BattleResult",
		40004: "WorldRaid_ReceiveReward",
		41000: "ResetableContent_Get",
		42000: "Conquest_GetInfo",
		42001: "Conquest_Conquer",
		42002: "Conquest_ConquerWithBattleStart",
		42003: "Conquest_ConquerWithBattleResult",
		42004: "Conquest_DeployEchelon",
		42005: "Conquest_ManageBase",
		42006: "Conquest_UpgradeBase",
		42007: "Conquest_TakeEventObject",
		42008: "Conquest_EventObjectBattleStart",
		42009: "Conquest_EventObjectBattleResult",
		42010: "Conquest_ReceiveCalculateRewards",
		42011: "Conquest_NormalizeEchelon",
		42012: "Conquest_Check",
		42013: "Conquest_ErosionBattleStart",
		42014: "Conquest_ErosionBattleResult",
		42015: "Conquest_MainStoryGetInfo",
		42016: "Conquest_MainStoryConquer",
		42017: "Conquest_MainStoryConquerWithBattleStart",
		42018: "Conquest_MainStoryConquerWithBattleResult",
		42019: "Conquest_MainStoryCheck",
		43000: "Friend_List",
		43001: "Friend_Remove",
		43002: "Friend_GetFriendDetailedInfo",
		43003: "Friend_GetIdCard",
		43004: "Friend_SetIdCard",
		43005: "Friend_Search",
		43006: "Friend_SendFriendRequest",
		43007: "Friend_AcceptFriendRequest",
		43008: "Friend_DeclineFriendRequest",
		43009: "Friend_CancelFriendRequest",
		43010: "Friend_Check",
		43011: "Friend_ListByIds",
		43012: "Friend_Block",
		43013: "Friend_Unblock",
		44000: "CharacterGear_List",
		44001: "CharacterGear_Unlock",
		44002: "CharacterGear_TierUp",
		45000: "EliminateRaid_Login",
		45001: "EliminateRaid_Lobby",
		45002: "EliminateRaid_OpponentList",
		45003: "EliminateRaid_GetBestTeam",
		45004: "EliminateRaid_CreateBattle",
		45005: "EliminateRaid_EnterBattle",
		45006: "EliminateRaid_EndBattle",
		45007: "EliminateRaid_GiveUp",
		45008: "EliminateRaid_Sweep",
		45009: "EliminateRaid_SeasonReward",
		45010: "EliminateRaid_RankingReward",
		45011: "EliminateRaid_LimitedReward",
		45012: "EliminateRaid_RankingIndex",
		46000: "Attachment_Get",
		46001: "Attachment_EmblemList",
		46002: "Attachment_EmblemAcquire",
		46003: "Attachment_EmblemAttach",
		47000: "Sticker_Login",
		47001: "Sticker_Lobby",
		47002: "Sticker_UseSticker",
		48000: "Field_Sync",
		48001: "Field_Interaction",
		48002: "Field_QuestClear",
		48003: "Field_SceneChanged",
		48004: "Field_EndDate",
		48005: "Field_EnterStage",
		48006: "Field_StageResult",
		49000: "MultiFloorRaid_Sync",
		49001: "MultiFloorRaid_EnterBattle",
		49002: "MultiFloorRaid_EndBattle",
		49003: "MultiFloorRaid_ReceiveReward",
		50000: "Queuing_GetTicket",
		51000: "BattlePass_GetInfo",
		51001: "BattlePass_BuyLevel",
		51002: "BattlePass_ReceiveReward",
		51003: "BattlePass_MissionList",
		51004: "BattlePass_MissionSingleReward",
		51005: "BattlePass_MissionMultipleReward",
		51006: "BattlePass_Check",
		90000: "GmTalk",
	}
	Protocol_value = map[string]int32{
		"Common_Cheat":          -9999,
		"Error":                 -1,
		"None":                  0,
		"System_Version":        1,
		"Session_Info":          2,
		"NetworkTime_Sync":      3,
		"NetworkTime_SyncReply": 4,
		"Audit_GachaStatistics": 5,
		"Account_Create":        1000,
		"Account_Nickname":      1001,
		"Account_Auth":          1002,
		"Account_CurrencySync":  1003,
		"Account_SetRepresentCharacterAndComment":   1004,
		"Account_GetTutorial":                       1005,
		"Account_SetTutorial":                       1006,
		"Account_PassCheck":                         1007,
		"Account_VerifyForYostar":                   1008,
		"Account_CheckYostar":                       1009,
		"Account_CallName":                          1010,
		"Account_BirthDay":                          1011,
		"Account_Auth2":                             1012,
		"Account_LinkReward":                        1013,
		"Account_ReportXignCodeCheater":             1014,
		"Account_DismissRepurchasablePopup":         1015,
		"Account_InvalidateToken":                   1016,
		"Account_LoginSync":                         1017,
		"Account_Reset":                             1018,
		"Account_RequestBirthdayMail":               1019,
		"Account_CheckAccountLevelReward":           1020,
		"Account_ReceiveAccountLevelReward":         1021,
		"Character_List":                            2000,
		"Character_Transcendence":                   2001,
		"Character_ExpGrowth":                       2002,
		"Character_FavorGrowth":                     2003,
		"Character_UpdateSkillLevel":                2004,
		"Character_UnlockWeapon":                    2005,
		"Character_WeaponExpGrowth":                 2006,
		"Character_WeaponTranscendence":             2007,
		"Character_SetFavorites":                    2008,
		"Character_SetCostume":                      2009,
		"Character_BatchSkillLevelUpdate":           2010,
		"Character_PotentialGrowth":                 2011,
		"Equipment_List":                            3000,
		"Equipment_Sell":                            3001,
		"Equipment_Equip":                           3002,
		"Equipment_LevelUp":                         3003,
		"Equipment_TierUp":                          3004,
		"Equipment_Lock":                            3005,
		"Equipment_BatchGrowth":                     3006,
		"Item_List":                                 4000,
		"Item_Sell":                                 4001,
		"Item_Consume":                              4002,
		"Item_Lock":                                 4003,
		"Item_BulkConsume":                          4004,
		"Item_SelectTicket":                         4005,
		"Item_AutoSynth":                            4006,
		"Echelon_List":                              5000,
		"Echelon_Save":                              5001,
		"Echelon_PresetList":                        5002,
		"Echelon_PresetSave":                        5003,
		"Echelon_PresetGroupRename":                 5004,
		"Campaign_List":                             6000,
		"Campaign_EnterMainStage":                   6001,
		"Campaign_ConfirmMainStage":                 6002,
		"Campaign_DeployEchelon":                    6003,
		"Campaign_WithdrawEchelon":                  6004,
		"Campaign_MapMove":                          6005,
		"Campaign_EndTurn":                          6006,
		"Campaign_EnterTactic":                      6007,
		"Campaign_TacticResult":                     6008,
		"Campaign_Retreat":                          6009,
		"Campaign_ChapterClearReward":               6010,
		"Campaign_Heal":                             6011,
		"Campaign_EnterSubStage":                    6012,
		"Campaign_SubStageResult":                   6013,
		"Campaign_Portal":                           6014,
		"Campaign_ConfirmTutorialStage":             6015,
		"Campaign_PurchasePlayCountHardStage":       6016,
		"Campaign_EnterTutorialStage":               6017,
		"Campaign_TutorialStageResult":              6018,
		"Campaign_RestartMainStage":                 6019,
		"Campaign_EnterMainStageStrategySkip":       6020,
		"Campaign_MainStageStrategySkipResult":      6021,
		"Mail_List":                                 7000,
		"Mail_Check":                                7001,
		"Mail_Receive":                              7002,
		"Mission_List":                              8000,
		"Mission_Reward":                            8001,
		"Mission_MultipleReward":                    8002,
		"Mission_GuideReward":                       8003,
		"Mission_MultipleGuideReward":               8004,
		"Mission_Sync":                              8005,
		"Mission_GuideMissionSeasonList":            8006,
		"Attendance_List":                           9000,
		"Attendance_Check":                          9001,
		"Attendance_Reward":                         9002,
		"Shop_BuyMerchandise":                       10000,
		"Shop_BuyGacha":                             10001,
		"Shop_List":                                 10002,
		"Shop_Refresh":                              10003,
		"Shop_BuyEligma":                            10004,
		"Shop_BuyGacha2":                            10005,
		"Shop_GachaRecruitList":                     10006,
		"Shop_BuyRefreshMerchandise":                10007,
		"Shop_BuyGacha3":                            10008,
		"Shop_BuyAP":                                10009,
		"Shop_BeforehandGachaGet":                   10010,
		"Shop_BeforehandGachaRun":                   10011,
		"Shop_BeforehandGachaSave":                  10012,
		"Shop_BeforehandGachaPick":                  10013,
		"Shop_PickupSelectionGachaGet":              10014,
		"Shop_PickupSelectionGachaSet":              10015,
		"Shop_PickupSelectionGachaBuy":              10016,
		"Recipe_Craft":                              11000,
		"MemoryLobby_List":                          12000,
		"MemoryLobby_SetMain":                       12001,
		"MemoryLobby_UpdateLobbyMode":               12002,
		"MemoryLobby_Interact":                      12003,
		"CumulativeTimeReward_List":                 13000,
		"CumulativeTimeReward_Reward":               13001,
		"OpenCondition_List":                        15000,
		"OpenCondition_Set":                         15001,
		"OpenCondition_EventList":                   15002,
		"Toast_List":                                16000,
		"Raid_List":                                 17000,
		"Raid_CompleteList":                         17001,
		"Raid_Detail":                               17002,
		"Raid_Search":                               17003,
		"Raid_CreateBattle":                         17004,
		"Raid_EnterBattle":                          17005,
		"Raid_BattleUpdate":                         17006,
		"Raid_EndBattle":                            17007,
		"Raid_Reward":                               17008,
		"Raid_RewardAll":                            17009,
		"Raid_Revive":                               17010,
		"Raid_Share":                                17011,
		"Raid_SeasonInfo":                           17012,
		"Raid_SeasonReward":                         17013,
		"Raid_Lobby":                                17014,
		"Raid_GiveUp":                               17015,
		"Raid_OpponentList":                         17016,
		"Raid_RankingReward":                        17017,
		"Raid_Login":                                17018,
		"Raid_Sweep":                                17019,
		"Raid_GetBestTeam":                          17020,
		"Raid_RankingIndex":                         17021,
		"SkipHistory_List":                          18000,
		"SkipHistory_Save":                          18001,
		"Scenario_List":                             19000,
		"Scenario_Clear":                            19001,
		"Scenario_GroupHistoryUpdate":               19002,
		"Scenario_Skip":                             19003,
		"Scenario_Select":                           19004,
		"Scenario_AccountStudentChange":             19005,
		"Scenario_LobbyStudentChange":               19006,
		"Scenario_SpecialLobbyChange":               19007,
		"Scenario_Enter":                            19008,
		"Scenario_EnterMainStage":                   19009,
		"Scenario_ConfirmMainStage":                 19010,
		"Scenario_DeployEchelon":                    19011,
		"Scenario_WithdrawEchelon":                  19012,
		"Scenario_MapMove":                          19013,
		"Scenario_EndTurn":                          19014,
		"Scenario_EnterTactic":                      19015,
		"Scenario_TacticResult":                     19016,
		"Scenario_Retreat":                          19017,
		"Scenario_Portal":                           19018,
		"Scenario_RestartMainStage":                 19019,
		"Scenario_SkipMainStage":                    19020,
		"Cafe_Get":                                  20000,
		"Cafe_Ack":                                  20001,
		"Cafe_Deploy":                               20002,
		"Cafe_Relocate":                             20003,
		"Cafe_Remove":                               20004,
		"Cafe_RemoveAll":                            20005,
		"Cafe_Interact":                             20006,
		"Cafe_ListPreset":                           20007,
		"Cafe_RenamePreset":                         20008,
		"Cafe_ClearPreset":                          20009,
		"Cafe_UpdatePresetFurniture":                20010,
		"Cafe_ApplyPreset":                          20011,
		"Cafe_RankUp":                               20012,
		"Cafe_ReceiveCurrency":                      20013,
		"Cafe_GiveGift":                             20014,
		"Cafe_SummonCharacter":                      20015,
		"Cafe_TrophyHistory":                        20016,
		"Cafe_ApplyTemplate":                        20017,
		"Cafe_Open":                                 20018,
		"Cafe_Travel":                               20019,
		"Craft_List":                                21000,
		"Craft_SelectNode":                          21001,
		"Craft_UpdateNodeLevel":                     21002,
		"Craft_BeginProcess":                        21003,
		"Craft_CompleteProcess":                     21004,
		"Craft_Reward":                              21005,
		"Craft_HistoryList":                         21006,
		"Craft_ShiftingBeginProcess":                21007,
		"Craft_ShiftingCompleteProcess":             21008,
		"Craft_ShiftingReward":                      21009,
		"Craft_AutoBeginProcess":                    21010,
		"Craft_CompleteProcessAll":                  21011,
		"Craft_RewardAll":                           21012,
		"Craft_ShiftingCompleteProcessAll":          21013,
		"Craft_ShiftingRewardAll":                   21014,
		"Arena_EnterLobby":                          22000,
		"Arena_Login":                               22001,
		"Arena_SettingChange":                       22002,
		"Arena_OpponentList":                        22003,
		"Arena_EnterBattle":                         22004,
		"Arena_EnterBattlePart1":                    22005,
		"Arena_EnterBattlePart2":                    22006,
		"Arena_BattleResult":                        22007,
		"Arena_CumulativeTimeReward":                22008,
		"Arena_DailyReward":                         22009,
		"Arena_RankList":                            22010,
		"Arena_History":                             22011,
		"Arena_RecordSync":                          22012,
		"Arena_TicketPurchase":                      22013,
		"Arena_DamageReport":                        22014,
		"Arena_CheckSeasonCloseReward":              22015,
		"Arena_SyncEchelonSettingTime":              22016,
		"WeekDungeon_List":                          23000,
		"WeekDungeon_EnterBattle":                   23001,
		"WeekDungeon_BattleResult":                  23002,
		"WeekDungeon_Retreat":                       23003,
		"Academy_GetInfo":                           24000,
		"Academy_AttendSchedule":                    24001,
		"Academy_AttendFavorSchedule":               24002,
		"Event_GetList":                             25000,
		"Event_GetImage":                            25001,
		"Event_UseCoupon":                           25002,
		"Event_RewardIncrease":                      25003,
		"ContentSave_Get":                           26000,
		"ContentSave_Discard":                       26001,
		"ContentSweep_Request":                      27000,
		"ContentSweep_MultiSweep":                   27001,
		"ContentSweep_MultiSweepPresetList":         27002,
		"ContentSweep_SetMultiSweepPreset":          27003,
		"ContentSweep_SetMultiSweepPresetName":      27004,
		"Clan_Lobby":                                28000,
		"Clan_Login":                                28001,
		"Clan_Search":                               28002,
		"Clan_Create":                               28003,
		"Clan_Member":                               28004,
		"Clan_Applicant":                            28005,
		"Clan_Join":                                 28006,
		"Clan_Quit":                                 28007,
		"Clan_Permit":                               28008,
		"Clan_Kick":                                 28009,
		"Clan_Setting":                              28010,
		"Clan_Confer":                               28011,
		"Clan_Dismiss":                              28012,
		"Clan_AutoJoin":                             28013,
		"Clan_MemberList":                           28014,
		"Clan_CancelApply":                          28015,
		"Clan_MyAssistList":                         28016,
		"Clan_SetAssist":                            28017,
		"Clan_ChatLog":                              28018,
		"Clan_Check":                                28019,
		"Clan_AllAssistList":                        28020,
		"Billing_TransactionStartByYostar":          29000,
		"Billing_TransactionEndByYostar":            29001,
		"Billing_PurchaseListByYostar":              29002,
		"Billing_PurchaseFreeProduct":               29003,
		"EventContent_AdventureList":                30000,
		"EventContent_EnterMainStage":               30001,
		"EventContent_ConfirmMainStage":             30002,
		"EventContent_EnterTactic":                  30003,
		"EventContent_TacticResult":                 30004,
		"EventContent_EnterSubStage":                30005,
		"EventContent_SubStageResult":               30006,
		"EventContent_DeployEchelon":                30007,
		"EventContent_WithdrawEchelon":              30008,
		"EventContent_MapMove":                      30009,
		"EventContent_EndTurn":                      30010,
		"EventContent_Retreat":                      30011,
		"EventContent_Portal":                       30012,
		"EventContent_PurchasePlayCountHardStage":   30013,
		"EventContent_ShopList":                     30014,
		"EventContent_ShopRefresh":                  30015,
		"EventContent_ReceiveStageTotalReward":      30016,
		"EventContent_EnterMainGroundStage":         30017,
		"EventContent_MainGroundStageResult":        30018,
		"EventContent_ShopBuyMerchandise":           30019,
		"EventContent_ShopBuyRefreshMerchandise":    30020,
		"EventContent_SelectBuff":                   30021,
		"EventContent_BoxGachaShopList":             30022,
		"EventContent_BoxGachaShopPurchase":         30023,
		"EventContent_BoxGachaShopRefresh":          30024,
		"EventContent_CollectionList":               30025,
		"EventContent_CollectionForMission":         30026,
		"EventContent_ScenarioGroupHistoryUpdate":   30027,
		"EventContent_CardShopList":                 30028,
		"EventContent_CardShopShuffle":              30029,
		"EventContent_CardShopPurchase":             30030,
		"EventContent_RestartMainStage":             30031,
		"EventContent_LocationGetInfo":              30032,
		"EventContent_LocationAttendSchedule":       30033,
		"EventContent_FortuneGachaPurchase":         30034,
		"EventContent_SubEventLobby":                30035,
		"EventContent_EnterStoryStage":              30036,
		"EventContent_StoryStageResult":             30037,
		"EventContent_DiceRaceLobby":                30038,
		"EventContent_DiceRaceRoll":                 30039,
		"EventContent_DiceRaceLapReward":            30040,
		"EventContent_PermanentList":                30041,
		"EventContent_DiceRaceUseItem":              30042,
		"EventContent_CardShopPurchaseAll":          30043,
		"EventContent_TreasureLobby":                30044,
		"EventContent_TreasureFlip":                 30045,
		"EventContent_TreasureNextRound":            30046,
		"TTS_GetFile":                               31000,
		"ContentLog_UIOpenStatistics":               32000,
		"MomoTalk_OutLine":                          33000,
		"MomoTalk_MessageList":                      33001,
		"MomoTalk_Read":                             33002,
		"MomoTalk_Reply":                            33003,
		"MomoTalk_FavorSchedule":                    33004,
		"ClearDeck_List":                            34000,
		"MiniGame_StageList":                        35000,
		"MiniGame_EnterStage":                       35001,
		"MiniGame_Result":                           35002,
		"MiniGame_MissionList":                      35003,
		"MiniGame_MissionReward":                    35004,
		"MiniGame_MissionMultipleReward":            35005,
		"MiniGame_ShootingLobby":                    35006,
		"MiniGame_ShootingBattleEnter":              35007,
		"MiniGame_ShootingBattleResult":             35008,
		"MiniGame_ShootingSweep":                    35009,
		"MiniGame_TableBoardSync":                   35010,
		"MiniGame_TableBoardMove":                   35011,
		"MiniGame_TableBoardEncounterInput":         35012,
		"MiniGame_TableBoardBattleEncounter":        35013,
		"MiniGame_TableBoardBattleRunAway":          35014,
		"MiniGame_TableBoardClearThema":             35015,
		"MiniGame_TableBoardUseItem":                35016,
		"MiniGame_TableBoardResurrect":              35017,
		"MiniGame_TableBoardSweep":                  35018,
		"MiniGame_TableBoardMoveThema":              35019,
		"MiniGame_DreamMakerGetInfo":                35020,
		"MiniGame_DreamMakerNewGame":                35021,
		"MiniGame_DreamMakerRestart":                35022,
		"MiniGame_DreamMakerAttendSchedule":         35023,
		"MiniGame_DreamMakerDailyClosing":           35024,
		"MiniGame_DreamMakerEnding":                 35025,
		"MiniGame_DefenseGetInfo":                   35026,
		"MiniGame_DefenseEnterBattle":               35027,
		"MiniGame_DefenseBattleResult":              35028,
		"MiniGame_RoadPuzzleGetInfo":                35029,
		"MiniGame_RoadPuzzleTilePlace":              35030,
		"MiniGame_RoadPuzzleSaveStage":              35031,
		"MiniGame_RoadPuzzleClearStage":             35032,
		"MiniGame_CCGLobby":                         35033,
		"MiniGame_CCGCreateGame":                    35034,
		"MiniGame_CCGSweep":                         35035,
		"MiniGame_CCGEnterStage":                    35036,
		"MiniGame_CCGEndStageDual":                  35037,
		"MiniGame_CCGEndStageEvent":                 35038,
		"MiniGame_CCGSelectRewardCard":              35039,
		"Minigame_CCGReplaceCharacter":              35040,
		"MiniGame_CCGSelectCampAction":              35041,
		"MiniGame_CCGCompleteGame":                  35042,
		"MiniGame_CCGGiveupGame":                    35043,
		"MiniGame_CCGRerollReward":                  35044,
		"MiniGame_CCGBuyPerk":                       35045,
		"Notification_LobbyCheck":                   36000,
		"Notification_EventContentReddotCheck":      36001,
		"ProofToken_RequestQuestion":                37000,
		"ProofToken_Submit":                         37001,
		"SchoolDungeon_List":                        38000,
		"SchoolDungeon_EnterBattle":                 38001,
		"SchoolDungeon_BattleResult":                38002,
		"SchoolDungeon_Retreat":                     38003,
		"TimeAttackDungeon_Lobby":                   39000,
		"TimeAttackDungeon_CreateBattle":            39001,
		"TimeAttackDungeon_EnterBattle":             39002,
		"TimeAttackDungeon_EndBattle":               39003,
		"TimeAttackDungeon_Sweep":                   39004,
		"TimeAttackDungeon_GiveUp":                  39005,
		"TimeAttackDungeon_Login":                   39006,
		"WorldRaid_Lobby":                           40000,
		"WorldRaid_BossList":                        40001,
		"WorldRaid_EnterBattle":                     40002,
		"WorldRaid_BattleResult":                    40003,
		"WorldRaid_ReceiveReward":                   40004,
		"ResetableContent_Get":                      41000,
		"Conquest_GetInfo":                          42000,
		"Conquest_Conquer":                          42001,
		"Conquest_ConquerWithBattleStart":           42002,
		"Conquest_ConquerWithBattleResult":          42003,
		"Conquest_DeployEchelon":                    42004,
		"Conquest_ManageBase":                       42005,
		"Conquest_UpgradeBase":                      42006,
		"Conquest_TakeEventObject":                  42007,
		"Conquest_EventObjectBattleStart":           42008,
		"Conquest_EventObjectBattleResult":          42009,
		"Conquest_ReceiveCalculateRewards":          42010,
		"Conquest_NormalizeEchelon":                 42011,
		"Conquest_Check":                            42012,
		"Conquest_ErosionBattleStart":               42013,
		"Conquest_ErosionBattleResult":              42014,
		"Conquest_MainStoryGetInfo":                 42015,
		"Conquest_MainStoryConquer":                 42016,
		"Conquest_MainStoryConquerWithBattleStart":  42017,
		"Conquest_MainStoryConquerWithBattleResult": 42018,
		"Conquest_MainStoryCheck":                   42019,
		"Friend_List":                               43000,
		"Friend_Remove":                             43001,
		"Friend_GetFriendDetailedInfo":              43002,
		"Friend_GetIdCard":                          43003,
		"Friend_SetIdCard":                          43004,
		"Friend_Search":                             43005,
		"Friend_SendFriendRequest":                  43006,
		"Friend_AcceptFriendRequest":                43007,
		"Friend_DeclineFriendRequest":               43008,
		"Friend_CancelFriendRequest":                43009,
		"Friend_Check":                              43010,
		"Friend_ListByIds":                          43011,
		"Friend_Block":                              43012,
		"Friend_Unblock":                            43013,
		"CharacterGear_List":                        44000,
		"CharacterGear_Unlock":                      44001,
		"CharacterGear_TierUp":                      44002,
		"EliminateRaid_Login":                       45000,
		"EliminateRaid_Lobby":                       45001,
		"EliminateRaid_OpponentList":                45002,
		"EliminateRaid_GetBestTeam":                 45003,
		"EliminateRaid_CreateBattle":                45004,
		"EliminateRaid_EnterBattle":                 45005,
		"EliminateRaid_EndBattle":                   45006,
		"EliminateRaid_GiveUp":                      45007,
		"EliminateRaid_Sweep":                       45008,
		"EliminateRaid_SeasonReward":                45009,
		"EliminateRaid_RankingReward":               45010,
		"EliminateRaid_LimitedReward":               45011,
		"EliminateRaid_RankingIndex":                45012,
		"Attachment_Get":                            46000,
		"Attachment_EmblemList":                     46001,
		"Attachment_EmblemAcquire":                  46002,
		"Attachment_EmblemAttach":                   46003,
		"Sticker_Login":                             47000,
		"Sticker_Lobby":                             47001,
		"Sticker_UseSticker":                        47002,
		"Field_Sync":                                48000,
		"Field_Interaction":                         48001,
		"Field_QuestClear":                          48002,
		"Field_SceneChanged":                        48003,
		"Field_EndDate":                             48004,
		"Field_EnterStage":                          48005,
		"Field_StageResult":                         48006,
		"MultiFloorRaid_Sync":                       49000,
		"MultiFloorRaid_EnterBattle":                49001,
		"MultiFloorRaid_EndBattle":                  49002,
		"MultiFloorRaid_ReceiveReward":              49003,
		"Queuing_GetTicket":                         50000,
		"BattlePass_GetInfo":                        51000,
		"BattlePass_BuyLevel":                       51001,
		"BattlePass_ReceiveReward":                  51002,
		"BattlePass_MissionList":                    51003,
		"BattlePass_MissionSingleReward":            51004,
		"BattlePass_MissionMultipleReward":          51005,
		"BattlePass_Check":                          51006,
		"GmTalk":                                    90000,
	}
)

func (x Protocol) String() string {
	return Protocol_name[int32(x)]
}

func (x Protocol) Value(sr string) Protocol {
	return Protocol(Protocol_value[sr])
}
