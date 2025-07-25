package proto

type WebAPIErrorCode int32

const (
	WebAPIErrorCode_None                                           WebAPIErrorCode = 0
	WebAPIErrorCode_InvalidPacket                                  WebAPIErrorCode = 1
	WebAPIErrorCode_InvalidProtocol                                WebAPIErrorCode = 2
	WebAPIErrorCode_InvalidSession                                 WebAPIErrorCode = 3
	WebAPIErrorCode_InvalidVersion                                 WebAPIErrorCode = 4
	WebAPIErrorCode_InternalServerError                            WebAPIErrorCode = 5
	WebAPIErrorCode_DBError                                        WebAPIErrorCode = 6
	WebAPIErrorCode_InvalidToken                                   WebAPIErrorCode = 7
	WebAPIErrorCode_FailedToLockAccount                            WebAPIErrorCode = 8
	WebAPIErrorCode_InvalidCheatError                              WebAPIErrorCode = 9
	WebAPIErrorCode_AccountCurrencyCannotAffordCost                WebAPIErrorCode = 10
	WebAPIErrorCode_ExceedTranscendenceCountLimit                  WebAPIErrorCode = 11
	WebAPIErrorCode_MailBoxFull                                    WebAPIErrorCode = 12
	WebAPIErrorCode_InventoryAlreadyFull                           WebAPIErrorCode = 13
	WebAPIErrorCode_AccountNotFound                                WebAPIErrorCode = 14
	WebAPIErrorCode_DataClassNotFound                              WebAPIErrorCode = 15
	WebAPIErrorCode_DataEntityNotFound                             WebAPIErrorCode = 16
	WebAPIErrorCode_AccountGemPaidCannotAffordCost                 WebAPIErrorCode = 17
	WebAPIErrorCode_AccountGemBonusCannotAffordCost                WebAPIErrorCode = 18
	WebAPIErrorCode_AccountItemCannotAffordCost                    WebAPIErrorCode = 19
	WebAPIErrorCode_APITimeoutError                                WebAPIErrorCode = 20
	WebAPIErrorCode_FunctionTimeoutError                           WebAPIErrorCode = 21
	WebAPIErrorCode_DBDistributeTransactionError                   WebAPIErrorCode = 22
	WebAPIErrorCode_OccasionalJobError                             WebAPIErrorCode = 23
	WebAPIErrorCode_FailedToConsumeParcel                          WebAPIErrorCode = 100
	WebAPIErrorCode_InvalidString                                  WebAPIErrorCode = 200
	WebAPIErrorCode_InvalidStringLength                            WebAPIErrorCode = 201
	WebAPIErrorCode_EmptyString                                    WebAPIErrorCode = 202
	WebAPIErrorCode_SpecialSymbolNotAllowed                        WebAPIErrorCode = 203
	WebAPIErrorCode_InvalidDate                                    WebAPIErrorCode = 300
	WebAPIErrorCode_CoolTimeRemain                                 WebAPIErrorCode = 301
	WebAPIErrorCode_TimeElapseError                                WebAPIErrorCode = 302
	WebAPIErrorCode_ClientSendBadRequest                           WebAPIErrorCode = 400
	WebAPIErrorCode_ClientSendTooManyRequest                       WebAPIErrorCode = 401
	WebAPIErrorCode_ClientSuspectedAsCheater                       WebAPIErrorCode = 402
	WebAPIErrorCode_CombatVerificationFailedInDev                  WebAPIErrorCode = 403
	WebAPIErrorCode_ServerFailedToHandleRequest                    WebAPIErrorCode = 500
	WebAPIErrorCode_DocumentDBFailedToHandleRequest                WebAPIErrorCode = 501
	WebAPIErrorCode_ServerCacheFailedToHandleRequest               WebAPIErrorCode = 502
	WebAPIErrorCode_ReconnectBundleUpdateRequired                  WebAPIErrorCode = 800
	WebAPIErrorCode_GatewayMakeStandbyNotSupport                   WebAPIErrorCode = 900
	WebAPIErrorCode_GatewayPassCheckNotSupport                     WebAPIErrorCode = 901
	WebAPIErrorCode_GatewayWaitingTicketTimeOut                    WebAPIErrorCode = 902
	WebAPIErrorCode_ClientUpdateRequire                            WebAPIErrorCode = 903
	WebAPIErrorCode_AccountCreateNoDevId                           WebAPIErrorCode = 1000
	WebAPIErrorCode_AccountCreateDuplicatedDevId                   WebAPIErrorCode = 1001
	WebAPIErrorCode_AccountAuthEmptyDevId                          WebAPIErrorCode = 1002
	WebAPIErrorCode_AccountAuthNotCreated                          WebAPIErrorCode = 1003
	WebAPIErrorCode_AccountAccessControlWithoutPermission          WebAPIErrorCode = 1004
	WebAPIErrorCode_AccountNicknameEmptyString                     WebAPIErrorCode = 1005
	WebAPIErrorCode_AccountNicknameSameName                        WebAPIErrorCode = 1006
	WebAPIErrorCode_AccountNicknameWithInvalidString               WebAPIErrorCode = 1007
	WebAPIErrorCode_AccountNicknameWithInvalidLength               WebAPIErrorCode = 1008
	WebAPIErrorCode_YostarServerNotSuccessStatusCode               WebAPIErrorCode = 1009
	WebAPIErrorCode_YostarNetworkException                         WebAPIErrorCode = 1010
	WebAPIErrorCode_YostarException                                WebAPIErrorCode = 1011
	WebAPIErrorCode_AccoountPassCheckNotSupportCheat               WebAPIErrorCode = 1012
	WebAPIErrorCode_AccountCreateFail                              WebAPIErrorCode = 1013
	WebAPIErrorCode_AccountAddPubliserAccountFail                  WebAPIErrorCode = 1014
	WebAPIErrorCode_AccountAddDevIdFail                            WebAPIErrorCode = 1015
	WebAPIErrorCode_AccountCreateAlreadyPublisherAccoundId         WebAPIErrorCode = 1016
	WebAPIErrorCode_AccountUpdateStateFail                         WebAPIErrorCode = 1017
	WebAPIErrorCode_YostarCheckFail                                WebAPIErrorCode = 1018
	WebAPIErrorCode_EnterTicketInvalid                             WebAPIErrorCode = 1019
	WebAPIErrorCode_EnterTicketTimeOut                             WebAPIErrorCode = 1020
	WebAPIErrorCode_EnterTicketUsed                                WebAPIErrorCode = 1021
	WebAPIErrorCode_AccountCommentLengthOverLimit                  WebAPIErrorCode = 1022
	WebAPIErrorCode_AccountUpdateBirthdayFailed                    WebAPIErrorCode = 1023
	WebAPIErrorCode_AccountLoginError                              WebAPIErrorCode = 1024
	WebAPIErrorCode_AccountCurrencySyncError                       WebAPIErrorCode = 1025
	WebAPIErrorCode_InvalidClientCookie                            WebAPIErrorCode = 1026
	WebAPIErrorCode_InappositeNicknameRestricted                   WebAPIErrorCode = 1027
	WebAPIErrorCode_InappositeCommentRestricted                    WebAPIErrorCode = 1028
	WebAPIErrorCode_InappositeCallnameRestricted                   WebAPIErrorCode = 1029
	WebAPIErrorCode_CharacterNotFound                              WebAPIErrorCode = 2000
	WebAPIErrorCode_CharacterLocked                                WebAPIErrorCode = 2001
	WebAPIErrorCode_CharacterAlreadyHas                            WebAPIErrorCode = 2002
	WebAPIErrorCode_CharacterAssignedEchelon                       WebAPIErrorCode = 2003
	WebAPIErrorCode_CharacterFavorDownException                    WebAPIErrorCode = 2004
	WebAPIErrorCode_CharacterFavorMaxLevelExceed                   WebAPIErrorCode = 2005
	WebAPIErrorCode_CannotLevelUpSkill                             WebAPIErrorCode = 2006
	WebAPIErrorCode_CharacterLevelAlreadyMax                       WebAPIErrorCode = 2007
	WebAPIErrorCode_InvalidCharacterExpGrowthRequest               WebAPIErrorCode = 2008
	WebAPIErrorCode_CharacterWeaponDataNotFound                    WebAPIErrorCode = 2009
	WebAPIErrorCode_CharacterWeaponNotFound                        WebAPIErrorCode = 2010
	WebAPIErrorCode_CharacterWeaponAlreadyUnlocked                 WebAPIErrorCode = 2011
	WebAPIErrorCode_CharacterWeaponUnlockConditionFail             WebAPIErrorCode = 2012
	WebAPIErrorCode_CharacterWeaponExpGrowthNotValidItem           WebAPIErrorCode = 2013
	WebAPIErrorCode_InvalidCharacterWeaponExpGrowthRequest         WebAPIErrorCode = 2014
	WebAPIErrorCode_CharacterWeaponTranscendenceRecipeNotFound     WebAPIErrorCode = 2015
	WebAPIErrorCode_CharacterWeaponTranscendenceConditionFail      WebAPIErrorCode = 2016
	WebAPIErrorCode_CharacterWeaponUpdateFail                      WebAPIErrorCode = 2017
	WebAPIErrorCode_CharacterGearNotFound                          WebAPIErrorCode = 2018
	WebAPIErrorCode_CharacterGearAlreadyEquiped                    WebAPIErrorCode = 2019
	WebAPIErrorCode_CharacterGearCannotTierUp                      WebAPIErrorCode = 2020
	WebAPIErrorCode_CharacterGearCannotUnlock                      WebAPIErrorCode = 2021
	WebAPIErrorCode_CharacterCostumeNotFound                       WebAPIErrorCode = 2022
	WebAPIErrorCode_CharacterCostumeAlreadySet                     WebAPIErrorCode = 2023
	WebAPIErrorCode_CharacterCannotEquipCostume                    WebAPIErrorCode = 2024
	WebAPIErrorCode_InvalidCharacterSkillLevelUpdateRequest        WebAPIErrorCode = 2025
	WebAPIErrorCode_InvalidCharacterPotentialGrowthRequest         WebAPIErrorCode = 2026
	WebAPIErrorCode_CharacterPotentialGrowthDataNotFound           WebAPIErrorCode = 2027
	WebAPIErrorCode_EquipmentNotFound                              WebAPIErrorCode = 3000
	WebAPIErrorCode_InvalidEquipmentExpGrowthRequest               WebAPIErrorCode = 3001
	WebAPIErrorCode_EquipmentNotMatchingSlotItemCategory           WebAPIErrorCode = 3002
	WebAPIErrorCode_EquipmentLocked                                WebAPIErrorCode = 3003
	WebAPIErrorCode_EquipmentAlreadyEquiped                        WebAPIErrorCode = 3004
	WebAPIErrorCode_EquipmentConsumeItemLimitCountOver             WebAPIErrorCode = 3005
	WebAPIErrorCode_EquipmentNotEquiped                            WebAPIErrorCode = 3006
	WebAPIErrorCode_EquipmentCanNotEquip                           WebAPIErrorCode = 3007
	WebAPIErrorCode_EquipmentIngredientEmtpy                       WebAPIErrorCode = 3008
	WebAPIErrorCode_EquipmentCannotLevelUp                         WebAPIErrorCode = 3009
	WebAPIErrorCode_EquipmentCannotTierUp                          WebAPIErrorCode = 3010
	WebAPIErrorCode_EquipmentGearCannotUnlock                      WebAPIErrorCode = 3011
	WebAPIErrorCode_EquipmentBatchGrowthNotValid                   WebAPIErrorCode = 3012
	WebAPIErrorCode_ItemNotFound                                   WebAPIErrorCode = 4000
	WebAPIErrorCode_ItemLocked                                     WebAPIErrorCode = 4001
	WebAPIErrorCode_ItemCreateWithoutStackCount                    WebAPIErrorCode = 4002
	WebAPIErrorCode_ItemCreateStackCountFull                       WebAPIErrorCode = 4003
	WebAPIErrorCode_ItemNotUsingType                               WebAPIErrorCode = 4004
	WebAPIErrorCode_ItemEnchantIngredientFail                      WebAPIErrorCode = 4005
	WebAPIErrorCode_ItemInvalidConsumeRequest                      WebAPIErrorCode = 4006
	WebAPIErrorCode_ItemInsufficientStackCount                     WebAPIErrorCode = 4007
	WebAPIErrorCode_ItemOverExpirationDateTime                     WebAPIErrorCode = 4008
	WebAPIErrorCode_ItemCannotAutoSynth                            WebAPIErrorCode = 4009
	WebAPIErrorCode_EchelonEmptyLeader                             WebAPIErrorCode = 5000
	WebAPIErrorCode_EchelonNotFound                                WebAPIErrorCode = 5001
	WebAPIErrorCode_EchelonNotDeployed                             WebAPIErrorCode = 5002
	WebAPIErrorCode_EchelonSlotOverMaxCount                        WebAPIErrorCode = 5003
	WebAPIErrorCode_EchelonAssignCharacterOnOtherEchelon           WebAPIErrorCode = 5004
	WebAPIErrorCode_EchelonTypeNotAcceptable                       WebAPIErrorCode = 5005
	WebAPIErrorCode_EchelonEmptyNotAcceptable                      WebAPIErrorCode = 5006
	WebAPIErrorCode_EchelonPresetInvalidSave                       WebAPIErrorCode = 5007
	WebAPIErrorCode_EchelonPresetLabelLengthInvalid                WebAPIErrorCode = 5008
	WebAPIErrorCode_CampaignStageNotOpen                           WebAPIErrorCode = 6000
	WebAPIErrorCode_CampaignStagePlayLimit                         WebAPIErrorCode = 6001
	WebAPIErrorCode_CampaignStageEnterFail                         WebAPIErrorCode = 6002
	WebAPIErrorCode_CampaignStageInvalidSaveData                   WebAPIErrorCode = 6003
	WebAPIErrorCode_CampaignStageNotPlayerTurn                     WebAPIErrorCode = 6004
	WebAPIErrorCode_CampaignStageStageNotFound                     WebAPIErrorCode = 6005
	WebAPIErrorCode_CampaignStageHistoryNotFound                   WebAPIErrorCode = 6006
	WebAPIErrorCode_CampaignStageChapterNotFound                   WebAPIErrorCode = 6007
	WebAPIErrorCode_CampaignStageEchelonNotFound                   WebAPIErrorCode = 6008
	WebAPIErrorCode_CampaignStageWithdrawedCannotReUse             WebAPIErrorCode = 6009
	WebAPIErrorCode_CampaignStageChapterRewardInvalidReward        WebAPIErrorCode = 6010
	WebAPIErrorCode_CampaignStageChapterRewardAlreadyReceived      WebAPIErrorCode = 6011
	WebAPIErrorCode_CampaignStageTacticWinnerInvalid               WebAPIErrorCode = 6012
	WebAPIErrorCode_CampaignStageActionCountZero                   WebAPIErrorCode = 6013
	WebAPIErrorCode_CampaignStageHealNotAcceptable                 WebAPIErrorCode = 6014
	WebAPIErrorCode_CampaignStageHealLimit                         WebAPIErrorCode = 6015
	WebAPIErrorCode_CampaignStageLocationCanNotEngage              WebAPIErrorCode = 6016
	WebAPIErrorCode_CampaignEncounterWaitingCannotEndTurn          WebAPIErrorCode = 6017
	WebAPIErrorCode_CampaignTacticResultEmpty                      WebAPIErrorCode = 6018
	WebAPIErrorCode_CampaignPortalExitNotFound                     WebAPIErrorCode = 6019
	WebAPIErrorCode_CampaignCannotReachDestination                 WebAPIErrorCode = 6020
	WebAPIErrorCode_CampaignChapterRewardConditionNotSatisfied     WebAPIErrorCode = 6021
	WebAPIErrorCode_CampaignStageDataInvalid                       WebAPIErrorCode = 6022
	WebAPIErrorCode_ContentSweepNotOpened                          WebAPIErrorCode = 6023
	WebAPIErrorCode_CampaignTacticSkipFailed                       WebAPIErrorCode = 6024
	WebAPIErrorCode_CampaignUnableToRemoveFixedEchelon             WebAPIErrorCode = 6025
	WebAPIErrorCode_CampaignCharacterIsNotWhitelist                WebAPIErrorCode = 6026
	WebAPIErrorCode_CampaignFailedToSkipStrategy                   WebAPIErrorCode = 6027
	WebAPIErrorCode_InvalidSweepRequest                            WebAPIErrorCode = 6028
	WebAPIErrorCode_MailReceiveRequestInvalid                      WebAPIErrorCode = 7000
	WebAPIErrorCode_MissionCannotComplete                          WebAPIErrorCode = 8000
	WebAPIErrorCode_MissionRewardInvalid                           WebAPIErrorCode = 8001
	WebAPIErrorCode_AttendanceInvalid                              WebAPIErrorCode = 9000
	WebAPIErrorCode_ShopExcelNotFound                              WebAPIErrorCode = 10000
	WebAPIErrorCode_ShopAndGoodsNotMatched                         WebAPIErrorCode = 10001
	WebAPIErrorCode_ShopGoodsNotFound                              WebAPIErrorCode = 10002
	WebAPIErrorCode_ShopExceedPurchaseCountLimit                   WebAPIErrorCode = 10003
	WebAPIErrorCode_ShopCannotRefresh                              WebAPIErrorCode = 10004
	WebAPIErrorCode_ShopInfoNotFound                               WebAPIErrorCode = 10005
	WebAPIErrorCode_ShopCannotPurchaseActionPointLimitOver         WebAPIErrorCode = 10006
	WebAPIErrorCode_ShopNotOpened                                  WebAPIErrorCode = 10007
	WebAPIErrorCode_ShopInvalidGoods                               WebAPIErrorCode = 10008
	WebAPIErrorCode_ShopInvalidCostOrReward                        WebAPIErrorCode = 10009
	WebAPIErrorCode_ShopEligmaOverPurchase                         WebAPIErrorCode = 10010
	WebAPIErrorCode_ShopFreeRecruitInvalid                         WebAPIErrorCode = 10011
	WebAPIErrorCode_ShopNewbieGachaInvalid                         WebAPIErrorCode = 10012
	WebAPIErrorCode_ShopCannotNewGoodsRefresh                      WebAPIErrorCode = 10013
	WebAPIErrorCode_GachaCostNotValid                              WebAPIErrorCode = 10014
	WebAPIErrorCode_ShopRestrictBuyWhenInventoryFull               WebAPIErrorCode = 10015
	WebAPIErrorCode_BeforehandGachaMetadataNotFound                WebAPIErrorCode = 10016
	WebAPIErrorCode_BeforehandGachaCandidateNotFound               WebAPIErrorCode = 10017
	WebAPIErrorCode_BeforehandGachaInvalidLastIndex                WebAPIErrorCode = 10018
	WebAPIErrorCode_BeforehandGachaInvalidSaveIndex                WebAPIErrorCode = 10019
	WebAPIErrorCode_BeforehandGachaInvalidPickIndex                WebAPIErrorCode = 10020
	WebAPIErrorCode_BeforehandGachaDuplicatedResults               WebAPIErrorCode = 10021
	WebAPIErrorCode_ShopCannotRefreshManually                      WebAPIErrorCode = 10022
	WebAPIErrorCode_PickupSelectionDataNotFound                    WebAPIErrorCode = 10023
	WebAPIErrorCode_PickupSelectionDataInvalid                     WebAPIErrorCode = 10024
	WebAPIErrorCode_RecipeCraftNoData                              WebAPIErrorCode = 11000
	WebAPIErrorCode_RecipeCraftInsufficientIngredients             WebAPIErrorCode = 11001
	WebAPIErrorCode_RecipeCraftDataError                           WebAPIErrorCode = 11002
	WebAPIErrorCode_MemoryLobbyNotFound                            WebAPIErrorCode = 12000
	WebAPIErrorCode_LobbyModeChangeFailed                          WebAPIErrorCode = 12001
	WebAPIErrorCode_CumulativeTimeRewardNotFound                   WebAPIErrorCode = 13000
	WebAPIErrorCode_CumulativeTimeRewardAlreadyReceipt             WebAPIErrorCode = 13001
	WebAPIErrorCode_CumulativeTimeRewardInsufficientConnectionTime WebAPIErrorCode = 13002
	WebAPIErrorCode_OpenConditionClosed                            WebAPIErrorCode = 14000
	WebAPIErrorCode_OpenConditionSetNotSupport                     WebAPIErrorCode = 14001
	WebAPIErrorCode_CafeNotFound                                   WebAPIErrorCode = 15000
	WebAPIErrorCode_CafeFurnitureNotFound                          WebAPIErrorCode = 15001
	WebAPIErrorCode_CafeDeployFail                                 WebAPIErrorCode = 15002
	WebAPIErrorCode_CafeRelocateFail                               WebAPIErrorCode = 15003
	WebAPIErrorCode_CafeInteractionNotFound                        WebAPIErrorCode = 15004
	WebAPIErrorCode_CafeProductionEmpty                            WebAPIErrorCode = 15005
	WebAPIErrorCode_CafeRankUpFail                                 WebAPIErrorCode = 15006
	WebAPIErrorCode_CafePresetNotFound                             WebAPIErrorCode = 15007
	WebAPIErrorCode_CafeRenamePresetFail                           WebAPIErrorCode = 15008
	WebAPIErrorCode_CafeClearPresetFail                            WebAPIErrorCode = 15009
	WebAPIErrorCode_CafeUpdatePresetFurnitureFail                  WebAPIErrorCode = 15010
	WebAPIErrorCode_CafeReservePresetActivationTimeFail            WebAPIErrorCode = 15011
	WebAPIErrorCode_CafePresetApplyFail                            WebAPIErrorCode = 15012
	WebAPIErrorCode_CafePresetIsEmpty                              WebAPIErrorCode = 15013
	WebAPIErrorCode_CafeAlreadyVisitCharacter                      WebAPIErrorCode = 15014
	WebAPIErrorCode_CafeCannotSummonCharacter                      WebAPIErrorCode = 15015
	WebAPIErrorCode_CafeCanRefreshVisitCharacter                   WebAPIErrorCode = 15016
	WebAPIErrorCode_CafeAlreadyInteraction                         WebAPIErrorCode = 15017
	WebAPIErrorCode_CafeTemplateNotFound                           WebAPIErrorCode = 15018
	WebAPIErrorCode_CafeAlreadyOpened                              WebAPIErrorCode = 15019
	WebAPIErrorCode_CafeNoPlaceToTravel                            WebAPIErrorCode = 15020
	WebAPIErrorCode_CafeCannotTravelToOwnCafe                      WebAPIErrorCode = 15021
	WebAPIErrorCode_CafeCannotTravel                               WebAPIErrorCode = 15022
	WebAPIErrorCode_CafeCannotTravel_CafeLock                      WebAPIErrorCode = 15023
	WebAPIErrorCode_ScenarioMode_Fail                              WebAPIErrorCode = 16000
	WebAPIErrorCode_ScenarioMode_DuplicatedScenarioModeId          WebAPIErrorCode = 16001
	WebAPIErrorCode_ScenarioMode_LimitClearedScenario              WebAPIErrorCode = 16002
	WebAPIErrorCode_ScenarioMode_LimitAccountLevel                 WebAPIErrorCode = 16003
	WebAPIErrorCode_ScenarioMode_LimitClearedStage                 WebAPIErrorCode = 16004
	WebAPIErrorCode_ScenarioMode_LimitClubStudent                  WebAPIErrorCode = 16005
	WebAPIErrorCode_ScenarioMode_FailInDBProcess                   WebAPIErrorCode = 16006
	WebAPIErrorCode_ScenarioGroup_DuplicatedScenarioGroupId        WebAPIErrorCode = 16007
	WebAPIErrorCode_ScenarioGroup_FailInDBProcess                  WebAPIErrorCode = 16008
	WebAPIErrorCode_ScenarioGroup_DataNotFound                     WebAPIErrorCode = 16009
	WebAPIErrorCode_ScenarioGroup_MeetupConditionFail              WebAPIErrorCode = 16010
	WebAPIErrorCode_CraftInfoNotFound                              WebAPIErrorCode = 17000
	WebAPIErrorCode_CraftCanNotCreateNode                          WebAPIErrorCode = 17001
	WebAPIErrorCode_CraftCanNotUpdateNode                          WebAPIErrorCode = 17002
	WebAPIErrorCode_CraftCanNotBeginProcess                        WebAPIErrorCode = 17003
	WebAPIErrorCode_CraftNodeDepthError                            WebAPIErrorCode = 17004
	WebAPIErrorCode_CraftAlreadyProcessing                         WebAPIErrorCode = 17005
	WebAPIErrorCode_CraftCanNotCompleteProcess                     WebAPIErrorCode = 17006
	WebAPIErrorCode_CraftProcessNotComplete                        WebAPIErrorCode = 17007
	WebAPIErrorCode_CraftInvalidIngredient                         WebAPIErrorCode = 17008
	WebAPIErrorCode_CraftError                                     WebAPIErrorCode = 17009
	WebAPIErrorCode_CraftInvalidData                               WebAPIErrorCode = 17010
	WebAPIErrorCode_CraftNotAvailableToCafePresets                 WebAPIErrorCode = 17011
	WebAPIErrorCode_CraftNotEnoughEmptySlotCount                   WebAPIErrorCode = 17012
	WebAPIErrorCode_CraftInvalidPresetSlotDB                       WebAPIErrorCode = 17013
	WebAPIErrorCode_RaidExcelDataNotFound                          WebAPIErrorCode = 18000
	WebAPIErrorCode_RaidSeasonNotOpen                              WebAPIErrorCode = 18001
	WebAPIErrorCode_RaidDBDataNotFound                             WebAPIErrorCode = 18002
	WebAPIErrorCode_RaidBattleNotFound                             WebAPIErrorCode = 18003
	WebAPIErrorCode_RaidBattleUpdateFail                           WebAPIErrorCode = 18004
	WebAPIErrorCode_RaidCompleteListEmpty                          WebAPIErrorCode = 18005
	WebAPIErrorCode_RaidRoomCanNotCreate                           WebAPIErrorCode = 18006
	WebAPIErrorCode_RaidActionPointZero                            WebAPIErrorCode = 18007
	WebAPIErrorCode_RaidTicketZero                                 WebAPIErrorCode = 18008
	WebAPIErrorCode_RaidRoomCanNotJoin                             WebAPIErrorCode = 18009
	WebAPIErrorCode_RaidRoomMaxPlayer                              WebAPIErrorCode = 18010
	WebAPIErrorCode_RaidRewardDataNotFound                         WebAPIErrorCode = 18011
	WebAPIErrorCode_RaidSeasonRewardNotFound                       WebAPIErrorCode = 18012
	WebAPIErrorCode_RaidSeasonAlreadyReceiveReward                 WebAPIErrorCode = 18013
	WebAPIErrorCode_RaidSeasonAddRewardPointError                  WebAPIErrorCode = 18014
	WebAPIErrorCode_RaidSeasonRewardNotUpdate                      WebAPIErrorCode = 18015
	WebAPIErrorCode_RaidSeasonReceiveRewardFail                    WebAPIErrorCode = 18016
	WebAPIErrorCode_RaidSearchNotFound                             WebAPIErrorCode = 18017
	WebAPIErrorCode_RaidShareNotFound                              WebAPIErrorCode = 18018
	WebAPIErrorCode_RaidEndRewardFlagError                         WebAPIErrorCode = 18019
	WebAPIErrorCode_RaidCanNotFoundPlayer                          WebAPIErrorCode = 18020
	WebAPIErrorCode_RaidAlreadyParticipateCharacters               WebAPIErrorCode = 18021
	WebAPIErrorCode_RaidClearHistoryNotSave                        WebAPIErrorCode = 18022
	WebAPIErrorCode_RaidBattleAlreadyEnd                           WebAPIErrorCode = 18023
	WebAPIErrorCode_RaidEchelonNotFound                            WebAPIErrorCode = 18024
	WebAPIErrorCode_RaidSeasonOpen                                 WebAPIErrorCode = 18025
	WebAPIErrorCode_RaidRoomIsAlreadyClose                         WebAPIErrorCode = 18026
	WebAPIErrorCode_RaidRankingNotFound                            WebAPIErrorCode = 18027
	WebAPIErrorCode_WeekDungeonInfoNotFound                        WebAPIErrorCode = 19000
	WebAPIErrorCode_WeekDungeonNotOpenToday                        WebAPIErrorCode = 19001
	WebAPIErrorCode_WeekDungeonBattleWinnerInvalid                 WebAPIErrorCode = 19002
	WebAPIErrorCode_WeekDungeonInvalidSaveData                     WebAPIErrorCode = 19003
	WebAPIErrorCode_FindGiftRewardNotFound                         WebAPIErrorCode = 20000
	WebAPIErrorCode_FindGiftRewardAlreadyAcquired                  WebAPIErrorCode = 20001
	WebAPIErrorCode_FindGiftClearCountOverTotalCount               WebAPIErrorCode = 20002
	WebAPIErrorCode_ArenaInfoNotFound                              WebAPIErrorCode = 21000
	WebAPIErrorCode_ArenaGroupNotFound                             WebAPIErrorCode = 21001
	WebAPIErrorCode_ArenaRankHistoryNotFound                       WebAPIErrorCode = 21002
	WebAPIErrorCode_ArenaRankInvalid                               WebAPIErrorCode = 21003
	WebAPIErrorCode_ArenaBattleFail                                WebAPIErrorCode = 21004
	WebAPIErrorCode_ArenaDailyRewardAlreadyBeenReceived            WebAPIErrorCode = 21005
	WebAPIErrorCode_ArenaNoSeasonAvailable                         WebAPIErrorCode = 21006
	WebAPIErrorCode_ArenaAttackCoolTime                            WebAPIErrorCode = 21007
	WebAPIErrorCode_ArenaOpponentAlreadyBeenAttacked               WebAPIErrorCode = 21008
	WebAPIErrorCode_ArenaOpponentRankInvalid                       WebAPIErrorCode = 21009
	WebAPIErrorCode_ArenaNeedFormationSetting                      WebAPIErrorCode = 21010
	WebAPIErrorCode_ArenaNoHistory                                 WebAPIErrorCode = 21011
	WebAPIErrorCode_ArenaInvalidRequest                            WebAPIErrorCode = 21012
	WebAPIErrorCode_ArenaInvalidIndex                              WebAPIErrorCode = 21013
	WebAPIErrorCode_ArenaNotFoundBattle                            WebAPIErrorCode = 21014
	WebAPIErrorCode_ArenaBattleTimeOver                            WebAPIErrorCode = 21015
	WebAPIErrorCode_ArenaRefreshTimeOver                           WebAPIErrorCode = 21016
	WebAPIErrorCode_ArenaEchelonSettingTimeOver                    WebAPIErrorCode = 21017
	WebAPIErrorCode_ArenaCannotReceiveReward                       WebAPIErrorCode = 21018
	WebAPIErrorCode_ArenaRewardNotExist                            WebAPIErrorCode = 21019
	WebAPIErrorCode_ArenaCannotSetMap                              WebAPIErrorCode = 21020
	WebAPIErrorCode_ArenaDefenderRankChange                        WebAPIErrorCode = 21021
	WebAPIErrorCode_AcademyNotFound                                WebAPIErrorCode = 22000
	WebAPIErrorCode_AcademyScheduleTableNotFound                   WebAPIErrorCode = 22001
	WebAPIErrorCode_AcademyScheduleOperationNotFound               WebAPIErrorCode = 22002
	WebAPIErrorCode_AcademyAlreadyAttendedSchedule                 WebAPIErrorCode = 22003
	WebAPIErrorCode_AcademyAlreadyAttendedFavorSchedule            WebAPIErrorCode = 22004
	WebAPIErrorCode_AcademyRewardCharacterNotFound                 WebAPIErrorCode = 22005
	WebAPIErrorCode_AcademyScheduleCanNotAttend                    WebAPIErrorCode = 22006
	WebAPIErrorCode_AcademyTicketZero                              WebAPIErrorCode = 22007
	WebAPIErrorCode_AcademyMessageCanNotSend                       WebAPIErrorCode = 22008
	WebAPIErrorCode_ContentSaveDBNotFound                          WebAPIErrorCode = 26000
	WebAPIErrorCode_ContentSaveDBEntranceFeeEmpty                  WebAPIErrorCode = 26001
	WebAPIErrorCode_AccountBanned                                  WebAPIErrorCode = 27000
	WebAPIErrorCode_ServerNowLoadingProhibitedWord                 WebAPIErrorCode = 28000
	WebAPIErrorCode_ServerIsUnderMaintenance                       WebAPIErrorCode = 28001
	WebAPIErrorCode_ServerMaintenanceSoon                          WebAPIErrorCode = 28002
	WebAPIErrorCode_AccountIsNotInWhiteList                        WebAPIErrorCode = 28003
	WebAPIErrorCode_ServerContentsLockUpdating                     WebAPIErrorCode = 28004
	WebAPIErrorCode_ServerContentsLock                             WebAPIErrorCode = 28005
	WebAPIErrorCode_CouponIsEmpty                                  WebAPIErrorCode = 29000
	WebAPIErrorCode_CouponIsInvalid                                WebAPIErrorCode = 29001
	WebAPIErrorCode_UseCouponUsedListReadFail                      WebAPIErrorCode = 29002
	WebAPIErrorCode_UseCouponUsedCoupon                            WebAPIErrorCode = 29003
	WebAPIErrorCode_UseCouponNotFoundSerials                       WebAPIErrorCode = 29004
	WebAPIErrorCode_UseCouponDeleteSerials                         WebAPIErrorCode = 29005
	WebAPIErrorCode_UseCouponUnapprovedSerials                     WebAPIErrorCode = 29006
	WebAPIErrorCode_UseCouponExpiredSerials                        WebAPIErrorCode = 29007
	WebAPIErrorCode_UseCouponMaximumSerials                        WebAPIErrorCode = 29008
	WebAPIErrorCode_UseCouponNotFoundMeta                          WebAPIErrorCode = 29009
	WebAPIErrorCode_UseCouponDuplicateUseCoupon                    WebAPIErrorCode = 29010
	WebAPIErrorCode_UseCouponDuplicateUseSerial                    WebAPIErrorCode = 29011
	WebAPIErrorCode_BillingStartShopCashIdNotFound                 WebAPIErrorCode = 30000
	WebAPIErrorCode_BillingStartNotServiceTime                     WebAPIErrorCode = 30001
	WebAPIErrorCode_BillingStartUseConditionCheckError             WebAPIErrorCode = 30002
	WebAPIErrorCode_BillingStartSmallLevel                         WebAPIErrorCode = 30003
	WebAPIErrorCode_BillingStartMaxPurchaseCount                   WebAPIErrorCode = 30004
	WebAPIErrorCode_BillingStartFailAddOrder                       WebAPIErrorCode = 30005
	WebAPIErrorCode_BillingStartExistPurchase                      WebAPIErrorCode = 30006
	WebAPIErrorCode_BillingEndFailGetOrder                         WebAPIErrorCode = 30007
	WebAPIErrorCode_BillingEndShopCashIdNotFound                   WebAPIErrorCode = 30008
	WebAPIErrorCode_BillingEndProductIdNotFound                    WebAPIErrorCode = 30009
	WebAPIErrorCode_BillingEndMonthlyProductIdNotFound             WebAPIErrorCode = 30010
	WebAPIErrorCode_BillingEndInvalidState                         WebAPIErrorCode = 30011
	WebAPIErrorCode_BillingEndFailUpdteState                       WebAPIErrorCode = 30012
	WebAPIErrorCode_BillingEndFailSendMail                         WebAPIErrorCode = 30013
	WebAPIErrorCode_BillingEndInvalidAccount                       WebAPIErrorCode = 30014
	WebAPIErrorCode_BillingEndNotFoundPurchaseCount                WebAPIErrorCode = 30015
	WebAPIErrorCode_BillingEndFailUpdteMonthlyProduct              WebAPIErrorCode = 30016
	WebAPIErrorCode_BillingStartMailFull                           WebAPIErrorCode = 30017
	WebAPIErrorCode_BillingStartInventoryAndMailFull               WebAPIErrorCode = 30018
	WebAPIErrorCode_BillingEndRecvedErrorMonthlyProduct            WebAPIErrorCode = 30019
	WebAPIErrorCode_MonthlyProductNotOutdated                      WebAPIErrorCode = 30020
	WebAPIErrorCode_BillingBattlePassProductNotExist               WebAPIErrorCode = 30021
	WebAPIErrorCode_BillingBattlePassInfo                          WebAPIErrorCode = 30022
	WebAPIErrorCode_BillingBattlePassInvalidBuyStep                WebAPIErrorCode = 30023
	WebAPIErrorCode_BillingNotFreeProduct                          WebAPIErrorCode = 30024
	WebAPIErrorCode_BillingPurchaseFreeProduct                     WebAPIErrorCode = 30025
	WebAPIErrorCode_BillingProductSelectionSlotEmpty               WebAPIErrorCode = 30026
	WebAPIErrorCode_BillingProductSelectionSlotNotMatch            WebAPIErrorCode = 30027
	WebAPIErrorCode_BillingProductSelectConditionFailed            WebAPIErrorCode = 30028
	WebAPIErrorCode_BillingProductSelectionSlotNotFound            WebAPIErrorCode = 30029
	WebAPIErrorCode_ClanNotFound                                   WebAPIErrorCode = 31000
	WebAPIErrorCode_ClanSearchFailed                               WebAPIErrorCode = 31001
	WebAPIErrorCode_ClanEmptySearchString                          WebAPIErrorCode = 31002
	WebAPIErrorCode_ClanAccountAlreadyJoinedClan                   WebAPIErrorCode = 31003
	WebAPIErrorCode_ClanAccountAlreadyQuitClan                     WebAPIErrorCode = 31004
	WebAPIErrorCode_ClanCreateFailed                               WebAPIErrorCode = 31005
	WebAPIErrorCode_ClanMemberExceedCapacity                       WebAPIErrorCode = 31006
	WebAPIErrorCode_ClanDoesNotHavePermission                      WebAPIErrorCode = 31007
	WebAPIErrorCode_ClanTargetAccountIsNotApplicant                WebAPIErrorCode = 31008
	WebAPIErrorCode_ClanMemberNotFound                             WebAPIErrorCode = 31009
	WebAPIErrorCode_ClanCanNotKick                                 WebAPIErrorCode = 31010
	WebAPIErrorCode_ClanCanNotDismiss                              WebAPIErrorCode = 31011
	WebAPIErrorCode_ClanCanNotQuit                                 WebAPIErrorCode = 31012
	WebAPIErrorCode_ClanRejoinCoolOff                              WebAPIErrorCode = 31013
	WebAPIErrorCode_ClanChangeMemberGradeFailed                    WebAPIErrorCode = 31014
	WebAPIErrorCode_ClanHasBeenDisMissed                           WebAPIErrorCode = 31015
	WebAPIErrorCode_ClanCannotChangeJoinOption                     WebAPIErrorCode = 31016
	WebAPIErrorCode_ClanExceedConferCountLimit                     WebAPIErrorCode = 31017
	WebAPIErrorCode_ClanBusy                                       WebAPIErrorCode = 31018
	WebAPIErrorCode_ClanNameEmptyString                            WebAPIErrorCode = 31019
	WebAPIErrorCode_ClanNameWithInvalidLength                      WebAPIErrorCode = 31020
	WebAPIErrorCode_ClanAssistCharacterAlreadyDeployed             WebAPIErrorCode = 31021
	WebAPIErrorCode_ClanAssistNotValidUse                          WebAPIErrorCode = 31022
	WebAPIErrorCode_ClanAssistCharacterChanged                     WebAPIErrorCode = 31023
	WebAPIErrorCode_ClanAssistCoolTime                             WebAPIErrorCode = 31024
	WebAPIErrorCode_ClanAssistAlreadyUsedInRaidRoom                WebAPIErrorCode = 31025
	WebAPIErrorCode_ClanAssistAlreadyUsedInTimeAttackDungeonRoom   WebAPIErrorCode = 31026
	WebAPIErrorCode_ClanAssistEchelonHasAssistOnly                 WebAPIErrorCode = 31027
	WebAPIErrorCode_PaymentInvalidSign                             WebAPIErrorCode = 32000
	WebAPIErrorCode_PaymentInvalidSeed1                            WebAPIErrorCode = 32001
	WebAPIErrorCode_PaymentInvalidSeed2                            WebAPIErrorCode = 32002
	WebAPIErrorCode_PaymentInvalidInput                            WebAPIErrorCode = 32003
	WebAPIErrorCode_PaymentNotFoundPurchase                        WebAPIErrorCode = 32004
	WebAPIErrorCode_PaymentGetPurchaseOrderNotZero                 WebAPIErrorCode = 32005
	WebAPIErrorCode_PaymentSetPurchaseOrderNotZero                 WebAPIErrorCode = 32006
	WebAPIErrorCode_PaymentException                               WebAPIErrorCode = 32007
	WebAPIErrorCode_PaymentInvalidState                            WebAPIErrorCode = 32008
	WebAPIErrorCode_SessionNotFound                                WebAPIErrorCode = 33000
	WebAPIErrorCode_SessionParseFail                               WebAPIErrorCode = 33001
	WebAPIErrorCode_SessionInvalidInput                            WebAPIErrorCode = 33002
	WebAPIErrorCode_SessionNotAuth                                 WebAPIErrorCode = 33003
	WebAPIErrorCode_SessionDuplicateLogin                          WebAPIErrorCode = 33004
	WebAPIErrorCode_SessionTimeOver                                WebAPIErrorCode = 33005
	WebAPIErrorCode_SessionInvalidVersion                          WebAPIErrorCode = 33006
	WebAPIErrorCode_SessionChangeDate                              WebAPIErrorCode = 33007
	WebAPIErrorCode_CallName_RenameCoolTime                        WebAPIErrorCode = 34000
	WebAPIErrorCode_CallName_EmptyString                           WebAPIErrorCode = 34001
	WebAPIErrorCode_CallName_InvalidString                         WebAPIErrorCode = 34002
	WebAPIErrorCode_CallName_TTSServerIsNotAvailable               WebAPIErrorCode = 34003
	WebAPIErrorCode_CouchbaseInvalidCas                            WebAPIErrorCode = 35000
	WebAPIErrorCode_CouchbaseOperationFailed                       WebAPIErrorCode = 35001
	WebAPIErrorCode_CouchbaseRollBackFailed                        WebAPIErrorCode = 35002
	WebAPIErrorCode_EventContentCannotSelectBuff                   WebAPIErrorCode = 36000
	WebAPIErrorCode_EventContentNoBuffGroupAvailable               WebAPIErrorCode = 36001
	WebAPIErrorCode_EventContentBuffGroupIdDuplicated              WebAPIErrorCode = 36002
	WebAPIErrorCode_EventContentNotOpen                            WebAPIErrorCode = 36003
	WebAPIErrorCode_EventContentNoTotalRewardAvailable             WebAPIErrorCode = 36004
	WebAPIErrorCode_EventContentBoxGachaPurchaseFailed             WebAPIErrorCode = 36005
	WebAPIErrorCode_EventContentBoxGachaCannotRefresh              WebAPIErrorCode = 36006
	WebAPIErrorCode_EventContentCardShopCannotShuffle              WebAPIErrorCode = 36007
	WebAPIErrorCode_EventContentElementDoesNotExist                WebAPIErrorCode = 36008
	WebAPIErrorCode_EventContentElementAlreadyPurchased            WebAPIErrorCode = 36009
	WebAPIErrorCode_EventContentLocationNotFound                   WebAPIErrorCode = 36010
	WebAPIErrorCode_EventContentLocationScheduleCanNotAttend       WebAPIErrorCode = 36011
	WebAPIErrorCode_EventContentDiceRaceDataNotFound               WebAPIErrorCode = 36012
	WebAPIErrorCode_EventContentDiceRaceAlreadyReceiveLapRewardAll WebAPIErrorCode = 36013
	WebAPIErrorCode_EventContentDiceRaceInvalidDiceRaceResultType  WebAPIErrorCode = 36014
	WebAPIErrorCode_EventContentTreasureDataNotFound               WebAPIErrorCode = 36015
	WebAPIErrorCode_EventContentTreasureNotComplete                WebAPIErrorCode = 36016
	WebAPIErrorCode_EventContentTreasureFlipFailed                 WebAPIErrorCode = 36017
	WebAPIErrorCode_MiniGameStageIsNotOpen                         WebAPIErrorCode = 37000
	WebAPIErrorCode_MiniGameStageInvalidResult                     WebAPIErrorCode = 37001
	WebAPIErrorCode_MiniGameShootingStageInvlid                    WebAPIErrorCode = 37002
	WebAPIErrorCode_MiniGameShootingCannotSweep                    WebAPIErrorCode = 37003
	WebAPIErrorCode_MiniGameTableBoardSaveNotExist                 WebAPIErrorCode = 37004
	WebAPIErrorCode_MiniGameTableBoardPlayerCannotMove             WebAPIErrorCode = 37005
	WebAPIErrorCode_MiniGameTableBoardNoActiveEncounter            WebAPIErrorCode = 37006
	WebAPIErrorCode_MiniGameTableBoardInvalidEncounterRequest      WebAPIErrorCode = 37007
	WebAPIErrorCode_MiniGameTableBoardProcessEncounterFailed       WebAPIErrorCode = 37008
	WebAPIErrorCode_MiniGameTableBoardItemNotExist                 WebAPIErrorCode = 37009
	WebAPIErrorCode_MiniGameTableBoardInvalidItemUse               WebAPIErrorCode = 37010
	WebAPIErrorCode_MiniGameTableBoardInvalidClearThemaRequest     WebAPIErrorCode = 37011
	WebAPIErrorCode_MiniGameTableBoardInvalidSeason                WebAPIErrorCode = 37012
	WebAPIErrorCode_MiniGameTableBoardInvalidResurrectRequest      WebAPIErrorCode = 37013
	WebAPIErrorCode_MiniGameTableBoardSweepConditionFail           WebAPIErrorCode = 37014
	WebAPIErrorCode_MiniGameTableBoardInvalidData                  WebAPIErrorCode = 37015
	WebAPIErrorCode_MiniGameDreamCannotStartNewGame                WebAPIErrorCode = 37016
	WebAPIErrorCode_MiniGameDreamCannotApplyMultiplier             WebAPIErrorCode = 37017
	WebAPIErrorCode_MiniGameDreamCannotReset                       WebAPIErrorCode = 37018
	WebAPIErrorCode_MiniGameDreamNotEnoughActionCount              WebAPIErrorCode = 37019
	WebAPIErrorCode_MiniGameDreamSaveNotExist                      WebAPIErrorCode = 37020
	WebAPIErrorCode_MiniGameDreamActionCountRemain                 WebAPIErrorCode = 37021
	WebAPIErrorCode_MiniGameDreamRoundNotComplete                  WebAPIErrorCode = 37022
	WebAPIErrorCode_MiniGameDreamRewardAlreadyReceived             WebAPIErrorCode = 37023
	WebAPIErrorCode_MiniGameDreamRoundCompleted                    WebAPIErrorCode = 37024
	WebAPIErrorCode_MiniGameShouldReceiveEndingReward              WebAPIErrorCode = 37025
	WebAPIErrorCode_MiniGameDefenseCannotUseCharacter              WebAPIErrorCode = 37026
	WebAPIErrorCode_MiniGameDefenseNotOpenStage                    WebAPIErrorCode = 37027
	WebAPIErrorCode_MiniGameDefenseCannotApplyMultiplier           WebAPIErrorCode = 37028
	WebAPIErrorCode_MiniGameRoadPuzzleInvalidTilePlacement         WebAPIErrorCode = 37029
	WebAPIErrorCode_MiniGameRoadPuzzleCannotTrainDeparture         WebAPIErrorCode = 37030
	WebAPIErrorCode_MiniGameRoadPuzzleAlreadyCleared               WebAPIErrorCode = 37031
	WebAPIErrorCode_MiniGameRoadPuzzleCannotSave                   WebAPIErrorCode = 37032
	WebAPIErrorCode_MiniGameCCGPlayingSaveAlreadyExists            WebAPIErrorCode = 37033
	WebAPIErrorCode_MiniGameCCGSaveNotExists                       WebAPIErrorCode = 37034
	WebAPIErrorCode_MiniGameCCGPlayingStageAlreadyExists           WebAPIErrorCode = 37035
	WebAPIErrorCode_MiniGameCCGPlayingStageNotExists               WebAPIErrorCode = 37036
	WebAPIErrorCode_MiniGameCCGInvalidOperation                    WebAPIErrorCode = 37037
	WebAPIErrorCode_MiniGameCCGSaveNotComplete                     WebAPIErrorCode = 37038
	WebAPIErrorCode_MiniGameCCGNoRerollPoint                       WebAPIErrorCode = 37039
	WebAPIErrorCode_ProofTokenNotSubmitted                         WebAPIErrorCode = 38000
	WebAPIErrorCode_SchoolDungeonInfoNotFound                      WebAPIErrorCode = 39000
	WebAPIErrorCode_SchoolDungeonNotOpened                         WebAPIErrorCode = 39001
	WebAPIErrorCode_SchoolDungeonInvalidSaveData                   WebAPIErrorCode = 39002
	WebAPIErrorCode_SchoolDungeonBattleWinnerInvalid               WebAPIErrorCode = 39003
	WebAPIErrorCode_SchoolDungeonInvalidReward                     WebAPIErrorCode = 39004
	WebAPIErrorCode_TimeAttackDungeonDataNotFound                  WebAPIErrorCode = 40000
	WebAPIErrorCode_TimeAttackDungeonNotOpen                       WebAPIErrorCode = 40001
	WebAPIErrorCode_TimeAttackDungeonRoomTimeOut                   WebAPIErrorCode = 40002
	WebAPIErrorCode_TimeAttackDungeonRoomPlayCountOver             WebAPIErrorCode = 40003
	WebAPIErrorCode_TimeAttackDungeonRoomAlreadyExists             WebAPIErrorCode = 40004
	WebAPIErrorCode_TimeAttackDungeonRoomAlreadyClosed             WebAPIErrorCode = 40005
	WebAPIErrorCode_TimeAttackDungeonRoomNotExist                  WebAPIErrorCode = 40006
	WebAPIErrorCode_TimeAttackDungeonInvalidRequest                WebAPIErrorCode = 40007
	WebAPIErrorCode_TimeAttackDungeonInvalidData                   WebAPIErrorCode = 40008
	WebAPIErrorCode_WorldRaidDataNotFound                          WebAPIErrorCode = 41000
	WebAPIErrorCode_WorldRaidSeasonNotOpen                         WebAPIErrorCode = 41001
	WebAPIErrorCode_WorldRaidBossGroupNotOpen                      WebAPIErrorCode = 41002
	WebAPIErrorCode_WorldRaidInvalidOpenCondition                  WebAPIErrorCode = 41003
	WebAPIErrorCode_WorldRaidDifficultyNotOpen                     WebAPIErrorCode = 41004
	WebAPIErrorCode_WorldRaidAssistCharacterLimitOver              WebAPIErrorCode = 41005
	WebAPIErrorCode_WorldRaidContainBlackListCharacter             WebAPIErrorCode = 41006
	WebAPIErrorCode_WorldRaidValidFixedEchelonSetting              WebAPIErrorCode = 41007
	WebAPIErrorCode_WorldRaidAlredayReceiveRewardAll               WebAPIErrorCode = 41008
	WebAPIErrorCode_WorldRaidCannotReceiveReward                   WebAPIErrorCode = 41009
	WebAPIErrorCode_WorldRaidBossAlreadyDead                       WebAPIErrorCode = 41010
	WebAPIErrorCode_WorldRaidNotAnotherBossKilled                  WebAPIErrorCode = 41011
	WebAPIErrorCode_WorldRaidBattleResultUpdateFailed              WebAPIErrorCode = 41012
	WebAPIErrorCode_WorldRaidGemEnterCountLimitOver                WebAPIErrorCode = 41013
	WebAPIErrorCode_WorldRaidCannotGemEnter                        WebAPIErrorCode = 41014
	WebAPIErrorCode_WorldRaidNeedClearScenarioBoss                 WebAPIErrorCode = 41015
	WebAPIErrorCode_WorldRaidBossIsAlive                           WebAPIErrorCode = 41016
	WebAPIErrorCode_ConquestDataNotFound                           WebAPIErrorCode = 42000
	WebAPIErrorCode_ConquestAlreadyConquested                      WebAPIErrorCode = 42001
	WebAPIErrorCode_ConquestNotFullyConquested                     WebAPIErrorCode = 42002
	WebAPIErrorCode_ConquestStepNotOpened                          WebAPIErrorCode = 42003
	WebAPIErrorCode_ConquestUnableToReach                          WebAPIErrorCode = 42004
	WebAPIErrorCode_ConquestUnableToAttack                         WebAPIErrorCode = 42005
	WebAPIErrorCode_ConquestEchelonChangedCountMax                 WebAPIErrorCode = 42006
	WebAPIErrorCode_ConquestEchelonNotFound                        WebAPIErrorCode = 42007
	WebAPIErrorCode_ConquestCharacterAlreadyDeployed               WebAPIErrorCode = 42008
	WebAPIErrorCode_ConquestMaxUpgrade                             WebAPIErrorCode = 42009
	WebAPIErrorCode_ConquestUnitNotFound                           WebAPIErrorCode = 42010
	WebAPIErrorCode_ConquestObjectNotFound                         WebAPIErrorCode = 42011
	WebAPIErrorCode_ConquestCalculateRewardNotFound                WebAPIErrorCode = 42012
	WebAPIErrorCode_ConquestInvalidTileType                        WebAPIErrorCode = 42013
	WebAPIErrorCode_ConquestInvalidObjectType                      WebAPIErrorCode = 42014
	WebAPIErrorCode_ConquestInvalidSaveData                        WebAPIErrorCode = 42015
	WebAPIErrorCode_ConquestMaxAssistCountReached                  WebAPIErrorCode = 42016
	WebAPIErrorCode_ConquestErosionConditionNotSatisfied           WebAPIErrorCode = 42017
	WebAPIErrorCode_ConquestAdditionalContentNotInUse              WebAPIErrorCode = 42018
	WebAPIErrorCode_ConquestCannotUseManageEchelon                 WebAPIErrorCode = 42019
	WebAPIErrorCode_FriendUserIsNotFriend                          WebAPIErrorCode = 43000
	WebAPIErrorCode_FriendFailedToCreateFriendIdCard               WebAPIErrorCode = 43001
	WebAPIErrorCode_FriendRequestNotFound                          WebAPIErrorCode = 43002
	WebAPIErrorCode_FriendInvalidFriendCode                        WebAPIErrorCode = 43003
	WebAPIErrorCode_FriendAlreadyFriend                            WebAPIErrorCode = 43004
	WebAPIErrorCode_FriendMaxSentRequestReached                    WebAPIErrorCode = 43005
	WebAPIErrorCode_FriendMaxReceivedRequestReached                WebAPIErrorCode = 43006
	WebAPIErrorCode_FriendCannotRequestMaxFriendCountReached       WebAPIErrorCode = 43007
	WebAPIErrorCode_FriendCannotAcceptMaxFriendCountReached        WebAPIErrorCode = 43008
	WebAPIErrorCode_FriendOpponentMaxFriendCountReached            WebAPIErrorCode = 43009
	WebAPIErrorCode_FriendTargetIsBusy                             WebAPIErrorCode = 43010
	WebAPIErrorCode_FriendRequestTargetIsYourself                  WebAPIErrorCode = 43011
	WebAPIErrorCode_FriendSearchTargetIsYourself                   WebAPIErrorCode = 43012
	WebAPIErrorCode_FriendInvalidBackgroundId                      WebAPIErrorCode = 43013
	WebAPIErrorCode_FriendIdCardCommentLengthOverLimit             WebAPIErrorCode = 43014
	WebAPIErrorCode_FriendBackgroundNotOwned                       WebAPIErrorCode = 43015
	WebAPIErrorCode_FriendBlockTargetIsYourself                    WebAPIErrorCode = 43016
	WebAPIErrorCode_FriendBlockTargetIsAlreadyBlocked              WebAPIErrorCode = 43017
	WebAPIErrorCode_FriendBlockTargetIsExceedMaxCount              WebAPIErrorCode = 43018
	WebAPIErrorCode_FriendBlockUserCannotOpenProfile               WebAPIErrorCode = 43019
	WebAPIErrorCode_FriendBlockUserCannotSendRequest               WebAPIErrorCode = 43020
	WebAPIErrorCode_EliminateStageIsNotOpened                      WebAPIErrorCode = 44000
	WebAPIErrorCode_MultiSweepPresetDocumentNotFound               WebAPIErrorCode = 45000
	WebAPIErrorCode_MultiSweepPresetNameEmpty                      WebAPIErrorCode = 45001
	WebAPIErrorCode_MultiSweepPresetInvalidStageId                 WebAPIErrorCode = 45002
	WebAPIErrorCode_MultiSweepPresetInvalidId                      WebAPIErrorCode = 45003
	WebAPIErrorCode_MultiSweepPresetNameInvalidLength              WebAPIErrorCode = 45004
	WebAPIErrorCode_MultiSweepPresetTooManySelectStageId           WebAPIErrorCode = 45005
	WebAPIErrorCode_MultiSweepPresetInvalidSweepCount              WebAPIErrorCode = 45006
	WebAPIErrorCode_MultiSweepPresetTooManySelectParcelId          WebAPIErrorCode = 45007
	WebAPIErrorCode_EmblemDataNotFound                             WebAPIErrorCode = 46000
	WebAPIErrorCode_EmblemAttachFailed                             WebAPIErrorCode = 46001
	WebAPIErrorCode_EmblemCannotReceive                            WebAPIErrorCode = 46002
	WebAPIErrorCode_EmblemPassCheckEmblemIsEmpty                   WebAPIErrorCode = 46003
	WebAPIErrorCode_StickerDataNotFound                            WebAPIErrorCode = 47000
	WebAPIErrorCode_StickerNotAcquired                             WebAPIErrorCode = 47001
	WebAPIErrorCode_StickerDocumentNotFound                        WebAPIErrorCode = 47002
	WebAPIErrorCode_StickerAlreadyUsed                             WebAPIErrorCode = 47003
	WebAPIErrorCode_ClearDeckInvalidKey                            WebAPIErrorCode = 48000
	WebAPIErrorCode_ClearDeckOutOfDate                             WebAPIErrorCode = 48001
	WebAPIErrorCode_FieldDataNotFound                              WebAPIErrorCode = 60000
	WebAPIErrorCode_FieldInteracionFailed                          WebAPIErrorCode = 60001
	WebAPIErrorCode_FieldQuestClearFailed                          WebAPIErrorCode = 60002
	WebAPIErrorCode_FieldInvalidSceneChangedRequest                WebAPIErrorCode = 60003
	WebAPIErrorCode_FieldInvalidEndDateRequest                     WebAPIErrorCode = 60004
	WebAPIErrorCode_FieldCreateDailyQuestFailed                    WebAPIErrorCode = 60005
	WebAPIErrorCode_FieldResetReplayFailed                         WebAPIErrorCode = 60006
	WebAPIErrorCode_FieldIncreaseMasteryFailed                     WebAPIErrorCode = 60007
	WebAPIErrorCode_FieldStageDataInvalid                          WebAPIErrorCode = 60008
	WebAPIErrorCode_FieldStageEnterFail                            WebAPIErrorCode = 60009
	WebAPIErrorCode_FieldContentIsClosed                           WebAPIErrorCode = 60010
	WebAPIErrorCode_FieldEventStageNotCleared                      WebAPIErrorCode = 60011
	WebAPIErrorCode_MultiFloorRaidSeasonNotOpened                  WebAPIErrorCode = 49000
	WebAPIErrorCode_MultiFloorRaidDataNotFound                     WebAPIErrorCode = 49001
	WebAPIErrorCode_MultiFloorRaidAssistCharacterLimitOver         WebAPIErrorCode = 49002
	WebAPIErrorCode_MultiFloorRaidStageOpenConditionFail           WebAPIErrorCode = 49003
	WebAPIErrorCode_MultiFloorRaidInvalidSummary                   WebAPIErrorCode = 49004
	WebAPIErrorCode_MultiFloorRaidInvalidRewardRequest             WebAPIErrorCode = 49005
	WebAPIErrorCode_BattlePassSeasonNotOpen                        WebAPIErrorCode = 50000
	WebAPIErrorCode_BattlePassBuyLevelAlreadyMaxLevel              WebAPIErrorCode = 50001
	WebAPIErrorCode_BattlePassBuyLevelMaxLevelOver                 WebAPIErrorCode = 50002
	WebAPIErrorCode_BattlePassBuyLevelBuyCountError                WebAPIErrorCode = 50003
	WebAPIErrorCode_BattlePassAlreadyGetRewardAll                  WebAPIErrorCode = 50004
)

var (
	WebAPIErrorCode_name = map[int32]string{
		0:     "None",
		1:     "InvalidPacket",
		2:     "InvalidProtocol",
		3:     "InvalidSession",
		4:     "InvalidVersion",
		5:     "InternalServerError",
		6:     "DBError",
		7:     "InvalidToken",
		8:     "FailedToLockAccount",
		9:     "InvalidCheatError",
		10:    "AccountCurrencyCannotAffordCost",
		11:    "ExceedTranscendenceCountLimit",
		12:    "MailBoxFull",
		13:    "InventoryAlreadyFull",
		14:    "AccountNotFound",
		15:    "DataClassNotFound",
		16:    "DataEntityNotFound",
		17:    "AccountGemPaidCannotAffordCost",
		18:    "AccountGemBonusCannotAffordCost",
		19:    "AccountItemCannotAffordCost",
		20:    "APITimeoutError",
		21:    "FunctionTimeoutError",
		22:    "DBDistributeTransactionError",
		23:    "OccasionalJobError",
		100:   "FailedToConsumeParcel",
		200:   "InvalidString",
		201:   "InvalidStringLength",
		202:   "EmptyString",
		203:   "SpecialSymbolNotAllowed",
		300:   "InvalidDate",
		301:   "CoolTimeRemain",
		302:   "TimeElapseError",
		400:   "ClientSendBadRequest",
		401:   "ClientSendTooManyRequest",
		402:   "ClientSuspectedAsCheater",
		403:   "CombatVerificationFailedInDev",
		500:   "ServerFailedToHandleRequest",
		501:   "DocumentDBFailedToHandleRequest",
		502:   "ServerCacheFailedToHandleRequest",
		800:   "ReconnectBundleUpdateRequired",
		900:   "GatewayMakeStandbyNotSupport",
		901:   "GatewayPassCheckNotSupport",
		902:   "GatewayWaitingTicketTimeOut",
		903:   "ClientUpdateRequire",
		1000:  "AccountCreateNoDevId",
		1001:  "AccountCreateDuplicatedDevId",
		1002:  "AccountAuthEmptyDevId",
		1003:  "AccountAuthNotCreated",
		1004:  "AccountAccessControlWithoutPermission",
		1005:  "AccountNicknameEmptyString",
		1006:  "AccountNicknameSameName",
		1007:  "AccountNicknameWithInvalidString",
		1008:  "AccountNicknameWithInvalidLength",
		1009:  "YostarServerNotSuccessStatusCode",
		1010:  "YostarNetworkException",
		1011:  "YostarException",
		1012:  "AccoountPassCheckNotSupportCheat",
		1013:  "AccountCreateFail",
		1014:  "AccountAddPubliserAccountFail",
		1015:  "AccountAddDevIdFail",
		1016:  "AccountCreateAlreadyPublisherAccoundId",
		1017:  "AccountUpdateStateFail",
		1018:  "YostarCheckFail",
		1019:  "EnterTicketInvalid",
		1020:  "EnterTicketTimeOut",
		1021:  "EnterTicketUsed",
		1022:  "AccountCommentLengthOverLimit",
		1023:  "AccountUpdateBirthdayFailed",
		1024:  "AccountLoginError",
		1025:  "AccountCurrencySyncError",
		1026:  "InvalidClientCookie",
		1027:  "InappositeNicknameRestricted",
		1028:  "InappositeCommentRestricted",
		1029:  "InappositeCallnameRestricted",
		2000:  "CharacterNotFound",
		2001:  "CharacterLocked",
		2002:  "CharacterAlreadyHas",
		2003:  "CharacterAssignedEchelon",
		2004:  "CharacterFavorDownException",
		2005:  "CharacterFavorMaxLevelExceed",
		2006:  "CannotLevelUpSkill",
		2007:  "CharacterLevelAlreadyMax",
		2008:  "InvalidCharacterExpGrowthRequest",
		2009:  "CharacterWeaponDataNotFound",
		2010:  "CharacterWeaponNotFound",
		2011:  "CharacterWeaponAlreadyUnlocked",
		2012:  "CharacterWeaponUnlockConditionFail",
		2013:  "CharacterWeaponExpGrowthNotValidItem",
		2014:  "InvalidCharacterWeaponExpGrowthRequest",
		2015:  "CharacterWeaponTranscendenceRecipeNotFound",
		2016:  "CharacterWeaponTranscendenceConditionFail",
		2017:  "CharacterWeaponUpdateFail",
		2018:  "CharacterGearNotFound",
		2019:  "CharacterGearAlreadyEquiped",
		2020:  "CharacterGearCannotTierUp",
		2021:  "CharacterGearCannotUnlock",
		2022:  "CharacterCostumeNotFound",
		2023:  "CharacterCostumeAlreadySet",
		2024:  "CharacterCannotEquipCostume",
		2025:  "InvalidCharacterSkillLevelUpdateRequest",
		2026:  "InvalidCharacterPotentialGrowthRequest",
		2027:  "CharacterPotentialGrowthDataNotFound",
		3000:  "EquipmentNotFound",
		3001:  "InvalidEquipmentExpGrowthRequest",
		3002:  "EquipmentNotMatchingSlotItemCategory",
		3003:  "EquipmentLocked",
		3004:  "EquipmentAlreadyEquiped",
		3005:  "EquipmentConsumeItemLimitCountOver",
		3006:  "EquipmentNotEquiped",
		3007:  "EquipmentCanNotEquip",
		3008:  "EquipmentIngredientEmtpy",
		3009:  "EquipmentCannotLevelUp",
		3010:  "EquipmentCannotTierUp",
		3011:  "EquipmentGearCannotUnlock",
		3012:  "EquipmentBatchGrowthNotValid",
		4000:  "ItemNotFound",
		4001:  "ItemLocked",
		4002:  "ItemCreateWithoutStackCount",
		4003:  "ItemCreateStackCountFull",
		4004:  "ItemNotUsingType",
		4005:  "ItemEnchantIngredientFail",
		4006:  "ItemInvalidConsumeRequest",
		4007:  "ItemInsufficientStackCount",
		4008:  "ItemOverExpirationDateTime",
		4009:  "ItemCannotAutoSynth",
		5000:  "EchelonEmptyLeader",
		5001:  "EchelonNotFound",
		5002:  "EchelonNotDeployed",
		5003:  "EchelonSlotOverMaxCount",
		5004:  "EchelonAssignCharacterOnOtherEchelon",
		5005:  "EchelonTypeNotAcceptable",
		5006:  "EchelonEmptyNotAcceptable",
		5007:  "EchelonPresetInvalidSave",
		5008:  "EchelonPresetLabelLengthInvalid",
		6000:  "CampaignStageNotOpen",
		6001:  "CampaignStagePlayLimit",
		6002:  "CampaignStageEnterFail",
		6003:  "CampaignStageInvalidSaveData",
		6004:  "CampaignStageNotPlayerTurn",
		6005:  "CampaignStageStageNotFound",
		6006:  "CampaignStageHistoryNotFound",
		6007:  "CampaignStageChapterNotFound",
		6008:  "CampaignStageEchelonNotFound",
		6009:  "CampaignStageWithdrawedCannotReUse",
		6010:  "CampaignStageChapterRewardInvalidReward",
		6011:  "CampaignStageChapterRewardAlreadyReceived",
		6012:  "CampaignStageTacticWinnerInvalid",
		6013:  "CampaignStageActionCountZero",
		6014:  "CampaignStageHealNotAcceptable",
		6015:  "CampaignStageHealLimit",
		6016:  "CampaignStageLocationCanNotEngage",
		6017:  "CampaignEncounterWaitingCannotEndTurn",
		6018:  "CampaignTacticResultEmpty",
		6019:  "CampaignPortalExitNotFound",
		6020:  "CampaignCannotReachDestination",
		6021:  "CampaignChapterRewardConditionNotSatisfied",
		6022:  "CampaignStageDataInvalid",
		6023:  "ContentSweepNotOpened",
		6024:  "CampaignTacticSkipFailed",
		6025:  "CampaignUnableToRemoveFixedEchelon",
		6026:  "CampaignCharacterIsNotWhitelist",
		6027:  "CampaignFailedToSkipStrategy",
		6028:  "InvalidSweepRequest",
		7000:  "MailReceiveRequestInvalid",
		8000:  "MissionCannotComplete",
		8001:  "MissionRewardInvalid",
		9000:  "AttendanceInvalid",
		10000: "ShopExcelNotFound",
		10001: "ShopAndGoodsNotMatched",
		10002: "ShopGoodsNotFound",
		10003: "ShopExceedPurchaseCountLimit",
		10004: "ShopCannotRefresh",
		10005: "ShopInfoNotFound",
		10006: "ShopCannotPurchaseActionPointLimitOver",
		10007: "ShopNotOpened",
		10008: "ShopInvalidGoods",
		10009: "ShopInvalidCostOrReward",
		10010: "ShopEligmaOverPurchase",
		10011: "ShopFreeRecruitInvalid",
		10012: "ShopNewbieGachaInvalid",
		10013: "ShopCannotNewGoodsRefresh",
		10014: "GachaCostNotValid",
		10015: "ShopRestrictBuyWhenInventoryFull",
		10016: "BeforehandGachaMetadataNotFound",
		10017: "BeforehandGachaCandidateNotFound",
		10018: "BeforehandGachaInvalidLastIndex",
		10019: "BeforehandGachaInvalidSaveIndex",
		10020: "BeforehandGachaInvalidPickIndex",
		10021: "BeforehandGachaDuplicatedResults",
		10022: "ShopCannotRefreshManually",
		10023: "PickupSelectionDataNotFound",
		10024: "PickupSelectionDataInvalid",
		11000: "RecipeCraftNoData",
		11001: "RecipeCraftInsufficientIngredients",
		11002: "RecipeCraftDataError",
		12000: "MemoryLobbyNotFound",
		12001: "LobbyModeChangeFailed",
		13000: "CumulativeTimeRewardNotFound",
		13001: "CumulativeTimeRewardAlreadyReceipt",
		13002: "CumulativeTimeRewardInsufficientConnectionTime",
		14000: "OpenConditionClosed",
		14001: "OpenConditionSetNotSupport",
		15000: "CafeNotFound",
		15001: "CafeFurnitureNotFound",
		15002: "CafeDeployFail",
		15003: "CafeRelocateFail",
		15004: "CafeInteractionNotFound",
		15005: "CafeProductionEmpty",
		15006: "CafeRankUpFail",
		15007: "CafePresetNotFound",
		15008: "CafeRenamePresetFail",
		15009: "CafeClearPresetFail",
		15010: "CafeUpdatePresetFurnitureFail",
		15011: "CafeReservePresetActivationTimeFail",
		15012: "CafePresetApplyFail",
		15013: "CafePresetIsEmpty",
		15014: "CafeAlreadyVisitCharacter",
		15015: "CafeCannotSummonCharacter",
		15016: "CafeCanRefreshVisitCharacter",
		15017: "CafeAlreadyInteraction",
		15018: "CafeTemplateNotFound",
		15019: "CafeAlreadyOpened",
		15020: "CafeNoPlaceToTravel",
		15021: "CafeCannotTravelToOwnCafe",
		15022: "CafeCannotTravel",
		15023: "CafeCannotTravel_CafeLock",
		16000: "ScenarioMode_Fail",
		16001: "ScenarioMode_DuplicatedScenarioModeId",
		16002: "ScenarioMode_LimitClearedScenario",
		16003: "ScenarioMode_LimitAccountLevel",
		16004: "ScenarioMode_LimitClearedStage",
		16005: "ScenarioMode_LimitClubStudent",
		16006: "ScenarioMode_FailInDBProcess",
		16007: "ScenarioGroup_DuplicatedScenarioGroupId",
		16008: "ScenarioGroup_FailInDBProcess",
		16009: "ScenarioGroup_DataNotFound",
		16010: "ScenarioGroup_MeetupConditionFail",
		17000: "CraftInfoNotFound",
		17001: "CraftCanNotCreateNode",
		17002: "CraftCanNotUpdateNode",
		17003: "CraftCanNotBeginProcess",
		17004: "CraftNodeDepthError",
		17005: "CraftAlreadyProcessing",
		17006: "CraftCanNotCompleteProcess",
		17007: "CraftProcessNotComplete",
		17008: "CraftInvalidIngredient",
		17009: "CraftError",
		17010: "CraftInvalidData",
		17011: "CraftNotAvailableToCafePresets",
		17012: "CraftNotEnoughEmptySlotCount",
		17013: "CraftInvalidPresetSlotDB",
		18000: "RaidExcelDataNotFound",
		18001: "RaidSeasonNotOpen",
		18002: "RaidDBDataNotFound",
		18003: "RaidBattleNotFound",
		18004: "RaidBattleUpdateFail",
		18005: "RaidCompleteListEmpty",
		18006: "RaidRoomCanNotCreate",
		18007: "RaidActionPointZero",
		18008: "RaidTicketZero",
		18009: "RaidRoomCanNotJoin",
		18010: "RaidRoomMaxPlayer",
		18011: "RaidRewardDataNotFound",
		18012: "RaidSeasonRewardNotFound",
		18013: "RaidSeasonAlreadyReceiveReward",
		18014: "RaidSeasonAddRewardPointError",
		18015: "RaidSeasonRewardNotUpdate",
		18016: "RaidSeasonReceiveRewardFail",
		18017: "RaidSearchNotFound",
		18018: "RaidShareNotFound",
		18019: "RaidEndRewardFlagError",
		18020: "RaidCanNotFoundPlayer",
		18021: "RaidAlreadyParticipateCharacters",
		18022: "RaidClearHistoryNotSave",
		18023: "RaidBattleAlreadyEnd",
		18024: "RaidEchelonNotFound",
		18025: "RaidSeasonOpen",
		18026: "RaidRoomIsAlreadyClose",
		18027: "RaidRankingNotFound",
		19000: "WeekDungeonInfoNotFound",
		19001: "WeekDungeonNotOpenToday",
		19002: "WeekDungeonBattleWinnerInvalid",
		19003: "WeekDungeonInvalidSaveData",
		20000: "FindGiftRewardNotFound",
		20001: "FindGiftRewardAlreadyAcquired",
		20002: "FindGiftClearCountOverTotalCount",
		21000: "ArenaInfoNotFound",
		21001: "ArenaGroupNotFound",
		21002: "ArenaRankHistoryNotFound",
		21003: "ArenaRankInvalid",
		21004: "ArenaBattleFail",
		21005: "ArenaDailyRewardAlreadyBeenReceived",
		21006: "ArenaNoSeasonAvailable",
		21007: "ArenaAttackCoolTime",
		21008: "ArenaOpponentAlreadyBeenAttacked",
		21009: "ArenaOpponentRankInvalid",
		21010: "ArenaNeedFormationSetting",
		21011: "ArenaNoHistory",
		21012: "ArenaInvalidRequest",
		21013: "ArenaInvalidIndex",
		21014: "ArenaNotFoundBattle",
		21015: "ArenaBattleTimeOver",
		21016: "ArenaRefreshTimeOver",
		21017: "ArenaEchelonSettingTimeOver",
		21018: "ArenaCannotReceiveReward",
		21019: "ArenaRewardNotExist",
		21020: "ArenaCannotSetMap",
		21021: "ArenaDefenderRankChange",
		22000: "AcademyNotFound",
		22001: "AcademyScheduleTableNotFound",
		22002: "AcademyScheduleOperationNotFound",
		22003: "AcademyAlreadyAttendedSchedule",
		22004: "AcademyAlreadyAttendedFavorSchedule",
		22005: "AcademyRewardCharacterNotFound",
		22006: "AcademyScheduleCanNotAttend",
		22007: "AcademyTicketZero",
		22008: "AcademyMessageCanNotSend",
		26000: "ContentSaveDBNotFound",
		26001: "ContentSaveDBEntranceFeeEmpty",
		27000: "AccountBanned",
		28000: "ServerNowLoadingProhibitedWord",
		28001: "ServerIsUnderMaintenance",
		28002: "ServerMaintenanceSoon",
		28003: "AccountIsNotInWhiteList",
		28004: "ServerContentsLockUpdating",
		28005: "ServerContentsLock",
		29000: "CouponIsEmpty",
		29001: "CouponIsInvalid",
		29002: "UseCouponUsedListReadFail",
		29003: "UseCouponUsedCoupon",
		29004: "UseCouponNotFoundSerials",
		29005: "UseCouponDeleteSerials",
		29006: "UseCouponUnapprovedSerials",
		29007: "UseCouponExpiredSerials",
		29008: "UseCouponMaximumSerials",
		29009: "UseCouponNotFoundMeta",
		29010: "UseCouponDuplicateUseCoupon",
		29011: "UseCouponDuplicateUseSerial",
		30000: "BillingStartShopCashIdNotFound",
		30001: "BillingStartNotServiceTime",
		30002: "BillingStartUseConditionCheckError",
		30003: "BillingStartSmallLevel",
		30004: "BillingStartMaxPurchaseCount",
		30005: "BillingStartFailAddOrder",
		30006: "BillingStartExistPurchase",
		30007: "BillingEndFailGetOrder",
		30008: "BillingEndShopCashIdNotFound",
		30009: "BillingEndProductIdNotFound",
		30010: "BillingEndMonthlyProductIdNotFound",
		30011: "BillingEndInvalidState",
		30012: "BillingEndFailUpdteState",
		30013: "BillingEndFailSendMail",
		30014: "BillingEndInvalidAccount",
		30015: "BillingEndNotFoundPurchaseCount",
		30016: "BillingEndFailUpdteMonthlyProduct",
		30017: "BillingStartMailFull",
		30018: "BillingStartInventoryAndMailFull",
		30019: "BillingEndRecvedErrorMonthlyProduct",
		30020: "MonthlyProductNotOutdated",
		30021: "BillingBattlePassProductNotExist",
		30022: "BillingBattlePassInfo",
		30023: "BillingBattlePassInvalidBuyStep",
		30024: "BillingNotFreeProduct",
		30025: "BillingPurchaseFreeProduct",
		30026: "BillingProductSelectionSlotEmpty",
		30027: "BillingProductSelectionSlotNotMatch",
		30028: "BillingProductSelectConditionFailed",
		30029: "BillingProductSelectionSlotNotFound",
		31000: "ClanNotFound",
		31001: "ClanSearchFailed",
		31002: "ClanEmptySearchString",
		31003: "ClanAccountAlreadyJoinedClan",
		31004: "ClanAccountAlreadyQuitClan",
		31005: "ClanCreateFailed",
		31006: "ClanMemberExceedCapacity",
		31007: "ClanDoesNotHavePermission",
		31008: "ClanTargetAccountIsNotApplicant",
		31009: "ClanMemberNotFound",
		31010: "ClanCanNotKick",
		31011: "ClanCanNotDismiss",
		31012: "ClanCanNotQuit",
		31013: "ClanRejoinCoolOff",
		31014: "ClanChangeMemberGradeFailed",
		31015: "ClanHasBeenDisMissed",
		31016: "ClanCannotChangeJoinOption",
		31017: "ClanExceedConferCountLimit",
		31018: "ClanBusy",
		31019: "ClanNameEmptyString",
		31020: "ClanNameWithInvalidLength",
		31021: "ClanAssistCharacterAlreadyDeployed",
		31022: "ClanAssistNotValidUse",
		31023: "ClanAssistCharacterChanged",
		31024: "ClanAssistCoolTime",
		31025: "ClanAssistAlreadyUsedInRaidRoom",
		31026: "ClanAssistAlreadyUsedInTimeAttackDungeonRoom",
		31027: "ClanAssistEchelonHasAssistOnly",
		32000: "PaymentInvalidSign",
		32001: "PaymentInvalidSeed1",
		32002: "PaymentInvalidSeed2",
		32003: "PaymentInvalidInput",
		32004: "PaymentNotFoundPurchase",
		32005: "PaymentGetPurchaseOrderNotZero",
		32006: "PaymentSetPurchaseOrderNotZero",
		32007: "PaymentException",
		32008: "PaymentInvalidState",
		33000: "SessionNotFound",
		33001: "SessionParseFail",
		33002: "SessionInvalidInput",
		33003: "SessionNotAuth",
		33004: "SessionDuplicateLogin",
		33005: "SessionTimeOver",
		33006: "SessionInvalidVersion",
		33007: "SessionChangeDate",
		34000: "CallName_RenameCoolTime",
		34001: "CallName_EmptyString",
		34002: "CallName_InvalidString",
		34003: "CallName_TTSServerIsNotAvailable",
		35000: "CouchbaseInvalidCas",
		35001: "CouchbaseOperationFailed",
		35002: "CouchbaseRollBackFailed",
		36000: "EventContentCannotSelectBuff",
		36001: "EventContentNoBuffGroupAvailable",
		36002: "EventContentBuffGroupIdDuplicated",
		36003: "EventContentNotOpen",
		36004: "EventContentNoTotalRewardAvailable",
		36005: "EventContentBoxGachaPurchaseFailed",
		36006: "EventContentBoxGachaCannotRefresh",
		36007: "EventContentCardShopCannotShuffle",
		36008: "EventContentElementDoesNotExist",
		36009: "EventContentElementAlreadyPurchased",
		36010: "EventContentLocationNotFound",
		36011: "EventContentLocationScheduleCanNotAttend",
		36012: "EventContentDiceRaceDataNotFound",
		36013: "EventContentDiceRaceAlreadyReceiveLapRewardAll",
		36014: "EventContentDiceRaceInvalidDiceRaceResultType",
		36015: "EventContentTreasureDataNotFound",
		36016: "EventContentTreasureNotComplete",
		36017: "EventContentTreasureFlipFailed",
		37000: "MiniGameStageIsNotOpen",
		37001: "MiniGameStageInvalidResult",
		37002: "MiniGameShootingStageInvlid",
		37003: "MiniGameShootingCannotSweep",
		37004: "MiniGameTableBoardSaveNotExist",
		37005: "MiniGameTableBoardPlayerCannotMove",
		37006: "MiniGameTableBoardNoActiveEncounter",
		37007: "MiniGameTableBoardInvalidEncounterRequest",
		37008: "MiniGameTableBoardProcessEncounterFailed",
		37009: "MiniGameTableBoardItemNotExist",
		37010: "MiniGameTableBoardInvalidItemUse",
		37011: "MiniGameTableBoardInvalidClearThemaRequest",
		37012: "MiniGameTableBoardInvalidSeason",
		37013: "MiniGameTableBoardInvalidResurrectRequest",
		37014: "MiniGameTableBoardSweepConditionFail",
		37015: "MiniGameTableBoardInvalidData",
		37016: "MiniGameDreamCannotStartNewGame",
		37017: "MiniGameDreamCannotApplyMultiplier",
		37018: "MiniGameDreamCannotReset",
		37019: "MiniGameDreamNotEnoughActionCount",
		37020: "MiniGameDreamSaveNotExist",
		37021: "MiniGameDreamActionCountRemain",
		37022: "MiniGameDreamRoundNotComplete",
		37023: "MiniGameDreamRewardAlreadyReceived",
		37024: "MiniGameDreamRoundCompleted",
		37025: "MiniGameShouldReceiveEndingReward",
		37026: "MiniGameDefenseCannotUseCharacter",
		37027: "MiniGameDefenseNotOpenStage",
		37028: "MiniGameDefenseCannotApplyMultiplier",
		37029: "MiniGameRoadPuzzleInvalidTilePlacement",
		37030: "MiniGameRoadPuzzleCannotTrainDeparture",
		37031: "MiniGameRoadPuzzleAlreadyCleared",
		37032: "MiniGameRoadPuzzleCannotSave",
		37033: "MiniGameCCGPlayingSaveAlreadyExists",
		37034: "MiniGameCCGSaveNotExists",
		37035: "MiniGameCCGPlayingStageAlreadyExists",
		37036: "MiniGameCCGPlayingStageNotExists",
		37037: "MiniGameCCGInvalidOperation",
		37038: "MiniGameCCGSaveNotComplete",
		37039: "MiniGameCCGNoRerollPoint",
		38000: "ProofTokenNotSubmitted",
		39000: "SchoolDungeonInfoNotFound",
		39001: "SchoolDungeonNotOpened",
		39002: "SchoolDungeonInvalidSaveData",
		39003: "SchoolDungeonBattleWinnerInvalid",
		39004: "SchoolDungeonInvalidReward",
		40000: "TimeAttackDungeonDataNotFound",
		40001: "TimeAttackDungeonNotOpen",
		40002: "TimeAttackDungeonRoomTimeOut",
		40003: "TimeAttackDungeonRoomPlayCountOver",
		40004: "TimeAttackDungeonRoomAlreadyExists",
		40005: "TimeAttackDungeonRoomAlreadyClosed",
		40006: "TimeAttackDungeonRoomNotExist",
		40007: "TimeAttackDungeonInvalidRequest",
		40008: "TimeAttackDungeonInvalidData",
		41000: "WorldRaidDataNotFound",
		41001: "WorldRaidSeasonNotOpen",
		41002: "WorldRaidBossGroupNotOpen",
		41003: "WorldRaidInvalidOpenCondition",
		41004: "WorldRaidDifficultyNotOpen",
		41005: "WorldRaidAssistCharacterLimitOver",
		41006: "WorldRaidContainBlackListCharacter",
		41007: "WorldRaidValidFixedEchelonSetting",
		41008: "WorldRaidAlredayReceiveRewardAll",
		41009: "WorldRaidCannotReceiveReward",
		41010: "WorldRaidBossAlreadyDead",
		41011: "WorldRaidNotAnotherBossKilled",
		41012: "WorldRaidBattleResultUpdateFailed",
		41013: "WorldRaidGemEnterCountLimitOver",
		41014: "WorldRaidCannotGemEnter",
		41015: "WorldRaidNeedClearScenarioBoss",
		41016: "WorldRaidBossIsAlive",
		42000: "ConquestDataNotFound",
		42001: "ConquestAlreadyConquested",
		42002: "ConquestNotFullyConquested",
		42003: "ConquestStepNotOpened",
		42004: "ConquestUnableToReach",
		42005: "ConquestUnableToAttack",
		42006: "ConquestEchelonChangedCountMax",
		42007: "ConquestEchelonNotFound",
		42008: "ConquestCharacterAlreadyDeployed",
		42009: "ConquestMaxUpgrade",
		42010: "ConquestUnitNotFound",
		42011: "ConquestObjectNotFound",
		42012: "ConquestCalculateRewardNotFound",
		42013: "ConquestInvalidTileType",
		42014: "ConquestInvalidObjectType",
		42015: "ConquestInvalidSaveData",
		42016: "ConquestMaxAssistCountReached",
		42017: "ConquestErosionConditionNotSatisfied",
		42018: "ConquestAdditionalContentNotInUse",
		42019: "ConquestCannotUseManageEchelon",
		43000: "FriendUserIsNotFriend",
		43001: "FriendFailedToCreateFriendIdCard",
		43002: "FriendRequestNotFound",
		43003: "FriendInvalidFriendCode",
		43004: "FriendAlreadyFriend",
		43005: "FriendMaxSentRequestReached",
		43006: "FriendMaxReceivedRequestReached",
		43007: "FriendCannotRequestMaxFriendCountReached",
		43008: "FriendCannotAcceptMaxFriendCountReached",
		43009: "FriendOpponentMaxFriendCountReached",
		43010: "FriendTargetIsBusy",
		43011: "FriendRequestTargetIsYourself",
		43012: "FriendSearchTargetIsYourself",
		43013: "FriendInvalidBackgroundId",
		43014: "FriendIdCardCommentLengthOverLimit",
		43015: "FriendBackgroundNotOwned",
		43016: "FriendBlockTargetIsYourself",
		43017: "FriendBlockTargetIsAlreadyBlocked",
		43018: "FriendBlockTargetIsExceedMaxCount",
		43019: "FriendBlockUserCannotOpenProfile",
		43020: "FriendBlockUserCannotSendRequest",
		44000: "EliminateStageIsNotOpened",
		45000: "MultiSweepPresetDocumentNotFound",
		45001: "MultiSweepPresetNameEmpty",
		45002: "MultiSweepPresetInvalidStageId",
		45003: "MultiSweepPresetInvalidId",
		45004: "MultiSweepPresetNameInvalidLength",
		45005: "MultiSweepPresetTooManySelectStageId",
		45006: "MultiSweepPresetInvalidSweepCount",
		45007: "MultiSweepPresetTooManySelectParcelId",
		46000: "EmblemDataNotFound",
		46001: "EmblemAttachFailed",
		46002: "EmblemCannotReceive",
		46003: "EmblemPassCheckEmblemIsEmpty",
		47000: "StickerDataNotFound",
		47001: "StickerNotAcquired",
		47002: "StickerDocumentNotFound",
		47003: "StickerAlreadyUsed",
		48000: "ClearDeckInvalidKey",
		48001: "ClearDeckOutOfDate",
		60000: "FieldDataNotFound",
		60001: "FieldInteracionFailed",
		60002: "FieldQuestClearFailed",
		60003: "FieldInvalidSceneChangedRequest",
		60004: "FieldInvalidEndDateRequest",
		60005: "FieldCreateDailyQuestFailed",
		60006: "FieldResetReplayFailed",
		60007: "FieldIncreaseMasteryFailed",
		60008: "FieldStageDataInvalid",
		60009: "FieldStageEnterFail",
		60010: "FieldContentIsClosed",
		60011: "FieldEventStageNotCleared",
		49000: "MultiFloorRaidSeasonNotOpened",
		49001: "MultiFloorRaidDataNotFound",
		49002: "MultiFloorRaidAssistCharacterLimitOver",
		49003: "MultiFloorRaidStageOpenConditionFail",
		49004: "MultiFloorRaidInvalidSummary",
		49005: "MultiFloorRaidInvalidRewardRequest",
		50000: "BattlePassSeasonNotOpen",
		50001: "BattlePassBuyLevelAlreadyMaxLevel",
		50002: "BattlePassBuyLevelMaxLevelOver",
		50003: "BattlePassBuyLevelBuyCountError",
		50004: "BattlePassAlreadyGetRewardAll",
	}
	WebAPIErrorCode_value = map[string]int32{
		"None":                                           0,
		"InvalidPacket":                                  1,
		"InvalidProtocol":                                2,
		"InvalidSession":                                 3,
		"InvalidVersion":                                 4,
		"InternalServerError":                            5,
		"DBError":                                        6,
		"InvalidToken":                                   7,
		"FailedToLockAccount":                            8,
		"InvalidCheatError":                              9,
		"AccountCurrencyCannotAffordCost":                10,
		"ExceedTranscendenceCountLimit":                  11,
		"MailBoxFull":                                    12,
		"InventoryAlreadyFull":                           13,
		"AccountNotFound":                                14,
		"DataClassNotFound":                              15,
		"DataEntityNotFound":                             16,
		"AccountGemPaidCannotAffordCost":                 17,
		"AccountGemBonusCannotAffordCost":                18,
		"AccountItemCannotAffordCost":                    19,
		"APITimeoutError":                                20,
		"FunctionTimeoutError":                           21,
		"DBDistributeTransactionError":                   22,
		"OccasionalJobError":                             23,
		"FailedToConsumeParcel":                          100,
		"InvalidString":                                  200,
		"InvalidStringLength":                            201,
		"EmptyString":                                    202,
		"SpecialSymbolNotAllowed":                        203,
		"InvalidDate":                                    300,
		"CoolTimeRemain":                                 301,
		"TimeElapseError":                                302,
		"ClientSendBadRequest":                           400,
		"ClientSendTooManyRequest":                       401,
		"ClientSuspectedAsCheater":                       402,
		"CombatVerificationFailedInDev":                  403,
		"ServerFailedToHandleRequest":                    500,
		"DocumentDBFailedToHandleRequest":                501,
		"ServerCacheFailedToHandleRequest":               502,
		"ReconnectBundleUpdateRequired":                  800,
		"GatewayMakeStandbyNotSupport":                   900,
		"GatewayPassCheckNotSupport":                     901,
		"GatewayWaitingTicketTimeOut":                    902,
		"ClientUpdateRequire":                            903,
		"AccountCreateNoDevId":                           1000,
		"AccountCreateDuplicatedDevId":                   1001,
		"AccountAuthEmptyDevId":                          1002,
		"AccountAuthNotCreated":                          1003,
		"AccountAccessControlWithoutPermission":          1004,
		"AccountNicknameEmptyString":                     1005,
		"AccountNicknameSameName":                        1006,
		"AccountNicknameWithInvalidString":               1007,
		"AccountNicknameWithInvalidLength":               1008,
		"YostarServerNotSuccessStatusCode":               1009,
		"YostarNetworkException":                         1010,
		"YostarException":                                1011,
		"AccoountPassCheckNotSupportCheat":               1012,
		"AccountCreateFail":                              1013,
		"AccountAddPubliserAccountFail":                  1014,
		"AccountAddDevIdFail":                            1015,
		"AccountCreateAlreadyPublisherAccoundId":         1016,
		"AccountUpdateStateFail":                         1017,
		"YostarCheckFail":                                1018,
		"EnterTicketInvalid":                             1019,
		"EnterTicketTimeOut":                             1020,
		"EnterTicketUsed":                                1021,
		"AccountCommentLengthOverLimit":                  1022,
		"AccountUpdateBirthdayFailed":                    1023,
		"AccountLoginError":                              1024,
		"AccountCurrencySyncError":                       1025,
		"InvalidClientCookie":                            1026,
		"InappositeNicknameRestricted":                   1027,
		"InappositeCommentRestricted":                    1028,
		"InappositeCallnameRestricted":                   1029,
		"CharacterNotFound":                              2000,
		"CharacterLocked":                                2001,
		"CharacterAlreadyHas":                            2002,
		"CharacterAssignedEchelon":                       2003,
		"CharacterFavorDownException":                    2004,
		"CharacterFavorMaxLevelExceed":                   2005,
		"CannotLevelUpSkill":                             2006,
		"CharacterLevelAlreadyMax":                       2007,
		"InvalidCharacterExpGrowthRequest":               2008,
		"CharacterWeaponDataNotFound":                    2009,
		"CharacterWeaponNotFound":                        2010,
		"CharacterWeaponAlreadyUnlocked":                 2011,
		"CharacterWeaponUnlockConditionFail":             2012,
		"CharacterWeaponExpGrowthNotValidItem":           2013,
		"InvalidCharacterWeaponExpGrowthRequest":         2014,
		"CharacterWeaponTranscendenceRecipeNotFound":     2015,
		"CharacterWeaponTranscendenceConditionFail":      2016,
		"CharacterWeaponUpdateFail":                      2017,
		"CharacterGearNotFound":                          2018,
		"CharacterGearAlreadyEquiped":                    2019,
		"CharacterGearCannotTierUp":                      2020,
		"CharacterGearCannotUnlock":                      2021,
		"CharacterCostumeNotFound":                       2022,
		"CharacterCostumeAlreadySet":                     2023,
		"CharacterCannotEquipCostume":                    2024,
		"InvalidCharacterSkillLevelUpdateRequest":        2025,
		"InvalidCharacterPotentialGrowthRequest":         2026,
		"CharacterPotentialGrowthDataNotFound":           2027,
		"EquipmentNotFound":                              3000,
		"InvalidEquipmentExpGrowthRequest":               3001,
		"EquipmentNotMatchingSlotItemCategory":           3002,
		"EquipmentLocked":                                3003,
		"EquipmentAlreadyEquiped":                        3004,
		"EquipmentConsumeItemLimitCountOver":             3005,
		"EquipmentNotEquiped":                            3006,
		"EquipmentCanNotEquip":                           3007,
		"EquipmentIngredientEmtpy":                       3008,
		"EquipmentCannotLevelUp":                         3009,
		"EquipmentCannotTierUp":                          3010,
		"EquipmentGearCannotUnlock":                      3011,
		"EquipmentBatchGrowthNotValid":                   3012,
		"ItemNotFound":                                   4000,
		"ItemLocked":                                     4001,
		"ItemCreateWithoutStackCount":                    4002,
		"ItemCreateStackCountFull":                       4003,
		"ItemNotUsingType":                               4004,
		"ItemEnchantIngredientFail":                      4005,
		"ItemInvalidConsumeRequest":                      4006,
		"ItemInsufficientStackCount":                     4007,
		"ItemOverExpirationDateTime":                     4008,
		"ItemCannotAutoSynth":                            4009,
		"EchelonEmptyLeader":                             5000,
		"EchelonNotFound":                                5001,
		"EchelonNotDeployed":                             5002,
		"EchelonSlotOverMaxCount":                        5003,
		"EchelonAssignCharacterOnOtherEchelon":           5004,
		"EchelonTypeNotAcceptable":                       5005,
		"EchelonEmptyNotAcceptable":                      5006,
		"EchelonPresetInvalidSave":                       5007,
		"EchelonPresetLabelLengthInvalid":                5008,
		"CampaignStageNotOpen":                           6000,
		"CampaignStagePlayLimit":                         6001,
		"CampaignStageEnterFail":                         6002,
		"CampaignStageInvalidSaveData":                   6003,
		"CampaignStageNotPlayerTurn":                     6004,
		"CampaignStageStageNotFound":                     6005,
		"CampaignStageHistoryNotFound":                   6006,
		"CampaignStageChapterNotFound":                   6007,
		"CampaignStageEchelonNotFound":                   6008,
		"CampaignStageWithdrawedCannotReUse":             6009,
		"CampaignStageChapterRewardInvalidReward":        6010,
		"CampaignStageChapterRewardAlreadyReceived":      6011,
		"CampaignStageTacticWinnerInvalid":               6012,
		"CampaignStageActionCountZero":                   6013,
		"CampaignStageHealNotAcceptable":                 6014,
		"CampaignStageHealLimit":                         6015,
		"CampaignStageLocationCanNotEngage":              6016,
		"CampaignEncounterWaitingCannotEndTurn":          6017,
		"CampaignTacticResultEmpty":                      6018,
		"CampaignPortalExitNotFound":                     6019,
		"CampaignCannotReachDestination":                 6020,
		"CampaignChapterRewardConditionNotSatisfied":     6021,
		"CampaignStageDataInvalid":                       6022,
		"ContentSweepNotOpened":                          6023,
		"CampaignTacticSkipFailed":                       6024,
		"CampaignUnableToRemoveFixedEchelon":             6025,
		"CampaignCharacterIsNotWhitelist":                6026,
		"CampaignFailedToSkipStrategy":                   6027,
		"InvalidSweepRequest":                            6028,
		"MailReceiveRequestInvalid":                      7000,
		"MissionCannotComplete":                          8000,
		"MissionRewardInvalid":                           8001,
		"AttendanceInvalid":                              9000,
		"ShopExcelNotFound":                              10000,
		"ShopAndGoodsNotMatched":                         10001,
		"ShopGoodsNotFound":                              10002,
		"ShopExceedPurchaseCountLimit":                   10003,
		"ShopCannotRefresh":                              10004,
		"ShopInfoNotFound":                               10005,
		"ShopCannotPurchaseActionPointLimitOver":         10006,
		"ShopNotOpened":                                  10007,
		"ShopInvalidGoods":                               10008,
		"ShopInvalidCostOrReward":                        10009,
		"ShopEligmaOverPurchase":                         10010,
		"ShopFreeRecruitInvalid":                         10011,
		"ShopNewbieGachaInvalid":                         10012,
		"ShopCannotNewGoodsRefresh":                      10013,
		"GachaCostNotValid":                              10014,
		"ShopRestrictBuyWhenInventoryFull":               10015,
		"BeforehandGachaMetadataNotFound":                10016,
		"BeforehandGachaCandidateNotFound":               10017,
		"BeforehandGachaInvalidLastIndex":                10018,
		"BeforehandGachaInvalidSaveIndex":                10019,
		"BeforehandGachaInvalidPickIndex":                10020,
		"BeforehandGachaDuplicatedResults":               10021,
		"ShopCannotRefreshManually":                      10022,
		"PickupSelectionDataNotFound":                    10023,
		"PickupSelectionDataInvalid":                     10024,
		"RecipeCraftNoData":                              11000,
		"RecipeCraftInsufficientIngredients":             11001,
		"RecipeCraftDataError":                           11002,
		"MemoryLobbyNotFound":                            12000,
		"LobbyModeChangeFailed":                          12001,
		"CumulativeTimeRewardNotFound":                   13000,
		"CumulativeTimeRewardAlreadyReceipt":             13001,
		"CumulativeTimeRewardInsufficientConnectionTime": 13002,
		"OpenConditionClosed":                            14000,
		"OpenConditionSetNotSupport":                     14001,
		"CafeNotFound":                                   15000,
		"CafeFurnitureNotFound":                          15001,
		"CafeDeployFail":                                 15002,
		"CafeRelocateFail":                               15003,
		"CafeInteractionNotFound":                        15004,
		"CafeProductionEmpty":                            15005,
		"CafeRankUpFail":                                 15006,
		"CafePresetNotFound":                             15007,
		"CafeRenamePresetFail":                           15008,
		"CafeClearPresetFail":                            15009,
		"CafeUpdatePresetFurnitureFail":                  15010,
		"CafeReservePresetActivationTimeFail":            15011,
		"CafePresetApplyFail":                            15012,
		"CafePresetIsEmpty":                              15013,
		"CafeAlreadyVisitCharacter":                      15014,
		"CafeCannotSummonCharacter":                      15015,
		"CafeCanRefreshVisitCharacter":                   15016,
		"CafeAlreadyInteraction":                         15017,
		"CafeTemplateNotFound":                           15018,
		"CafeAlreadyOpened":                              15019,
		"CafeNoPlaceToTravel":                            15020,
		"CafeCannotTravelToOwnCafe":                      15021,
		"CafeCannotTravel":                               15022,
		"CafeCannotTravel_CafeLock":                      15023,
		"ScenarioMode_Fail":                              16000,
		"ScenarioMode_DuplicatedScenarioModeId":          16001,
		"ScenarioMode_LimitClearedScenario":              16002,
		"ScenarioMode_LimitAccountLevel":                 16003,
		"ScenarioMode_LimitClearedStage":                 16004,
		"ScenarioMode_LimitClubStudent":                  16005,
		"ScenarioMode_FailInDBProcess":                   16006,
		"ScenarioGroup_DuplicatedScenarioGroupId":        16007,
		"ScenarioGroup_FailInDBProcess":                  16008,
		"ScenarioGroup_DataNotFound":                     16009,
		"ScenarioGroup_MeetupConditionFail":              16010,
		"CraftInfoNotFound":                              17000,
		"CraftCanNotCreateNode":                          17001,
		"CraftCanNotUpdateNode":                          17002,
		"CraftCanNotBeginProcess":                        17003,
		"CraftNodeDepthError":                            17004,
		"CraftAlreadyProcessing":                         17005,
		"CraftCanNotCompleteProcess":                     17006,
		"CraftProcessNotComplete":                        17007,
		"CraftInvalidIngredient":                         17008,
		"CraftError":                                     17009,
		"CraftInvalidData":                               17010,
		"CraftNotAvailableToCafePresets":                 17011,
		"CraftNotEnoughEmptySlotCount":                   17012,
		"CraftInvalidPresetSlotDB":                       17013,
		"RaidExcelDataNotFound":                          18000,
		"RaidSeasonNotOpen":                              18001,
		"RaidDBDataNotFound":                             18002,
		"RaidBattleNotFound":                             18003,
		"RaidBattleUpdateFail":                           18004,
		"RaidCompleteListEmpty":                          18005,
		"RaidRoomCanNotCreate":                           18006,
		"RaidActionPointZero":                            18007,
		"RaidTicketZero":                                 18008,
		"RaidRoomCanNotJoin":                             18009,
		"RaidRoomMaxPlayer":                              18010,
		"RaidRewardDataNotFound":                         18011,
		"RaidSeasonRewardNotFound":                       18012,
		"RaidSeasonAlreadyReceiveReward":                 18013,
		"RaidSeasonAddRewardPointError":                  18014,
		"RaidSeasonRewardNotUpdate":                      18015,
		"RaidSeasonReceiveRewardFail":                    18016,
		"RaidSearchNotFound":                             18017,
		"RaidShareNotFound":                              18018,
		"RaidEndRewardFlagError":                         18019,
		"RaidCanNotFoundPlayer":                          18020,
		"RaidAlreadyParticipateCharacters":               18021,
		"RaidClearHistoryNotSave":                        18022,
		"RaidBattleAlreadyEnd":                           18023,
		"RaidEchelonNotFound":                            18024,
		"RaidSeasonOpen":                                 18025,
		"RaidRoomIsAlreadyClose":                         18026,
		"RaidRankingNotFound":                            18027,
		"WeekDungeonInfoNotFound":                        19000,
		"WeekDungeonNotOpenToday":                        19001,
		"WeekDungeonBattleWinnerInvalid":                 19002,
		"WeekDungeonInvalidSaveData":                     19003,
		"FindGiftRewardNotFound":                         20000,
		"FindGiftRewardAlreadyAcquired":                  20001,
		"FindGiftClearCountOverTotalCount":               20002,
		"ArenaInfoNotFound":                              21000,
		"ArenaGroupNotFound":                             21001,
		"ArenaRankHistoryNotFound":                       21002,
		"ArenaRankInvalid":                               21003,
		"ArenaBattleFail":                                21004,
		"ArenaDailyRewardAlreadyBeenReceived":            21005,
		"ArenaNoSeasonAvailable":                         21006,
		"ArenaAttackCoolTime":                            21007,
		"ArenaOpponentAlreadyBeenAttacked":               21008,
		"ArenaOpponentRankInvalid":                       21009,
		"ArenaNeedFormationSetting":                      21010,
		"ArenaNoHistory":                                 21011,
		"ArenaInvalidRequest":                            21012,
		"ArenaInvalidIndex":                              21013,
		"ArenaNotFoundBattle":                            21014,
		"ArenaBattleTimeOver":                            21015,
		"ArenaRefreshTimeOver":                           21016,
		"ArenaEchelonSettingTimeOver":                    21017,
		"ArenaCannotReceiveReward":                       21018,
		"ArenaRewardNotExist":                            21019,
		"ArenaCannotSetMap":                              21020,
		"ArenaDefenderRankChange":                        21021,
		"AcademyNotFound":                                22000,
		"AcademyScheduleTableNotFound":                   22001,
		"AcademyScheduleOperationNotFound":               22002,
		"AcademyAlreadyAttendedSchedule":                 22003,
		"AcademyAlreadyAttendedFavorSchedule":            22004,
		"AcademyRewardCharacterNotFound":                 22005,
		"AcademyScheduleCanNotAttend":                    22006,
		"AcademyTicketZero":                              22007,
		"AcademyMessageCanNotSend":                       22008,
		"ContentSaveDBNotFound":                          26000,
		"ContentSaveDBEntranceFeeEmpty":                  26001,
		"AccountBanned":                                  27000,
		"ServerNowLoadingProhibitedWord":                 28000,
		"ServerIsUnderMaintenance":                       28001,
		"ServerMaintenanceSoon":                          28002,
		"AccountIsNotInWhiteList":                        28003,
		"ServerContentsLockUpdating":                     28004,
		"ServerContentsLock":                             28005,
		"CouponIsEmpty":                                  29000,
		"CouponIsInvalid":                                29001,
		"UseCouponUsedListReadFail":                      29002,
		"UseCouponUsedCoupon":                            29003,
		"UseCouponNotFoundSerials":                       29004,
		"UseCouponDeleteSerials":                         29005,
		"UseCouponUnapprovedSerials":                     29006,
		"UseCouponExpiredSerials":                        29007,
		"UseCouponMaximumSerials":                        29008,
		"UseCouponNotFoundMeta":                          29009,
		"UseCouponDuplicateUseCoupon":                    29010,
		"UseCouponDuplicateUseSerial":                    29011,
		"BillingStartShopCashIdNotFound":                 30000,
		"BillingStartNotServiceTime":                     30001,
		"BillingStartUseConditionCheckError":             30002,
		"BillingStartSmallLevel":                         30003,
		"BillingStartMaxPurchaseCount":                   30004,
		"BillingStartFailAddOrder":                       30005,
		"BillingStartExistPurchase":                      30006,
		"BillingEndFailGetOrder":                         30007,
		"BillingEndShopCashIdNotFound":                   30008,
		"BillingEndProductIdNotFound":                    30009,
		"BillingEndMonthlyProductIdNotFound":             30010,
		"BillingEndInvalidState":                         30011,
		"BillingEndFailUpdteState":                       30012,
		"BillingEndFailSendMail":                         30013,
		"BillingEndInvalidAccount":                       30014,
		"BillingEndNotFoundPurchaseCount":                30015,
		"BillingEndFailUpdteMonthlyProduct":              30016,
		"BillingStartMailFull":                           30017,
		"BillingStartInventoryAndMailFull":               30018,
		"BillingEndRecvedErrorMonthlyProduct":            30019,
		"MonthlyProductNotOutdated":                      30020,
		"BillingBattlePassProductNotExist":               30021,
		"BillingBattlePassInfo":                          30022,
		"BillingBattlePassInvalidBuyStep":                30023,
		"BillingNotFreeProduct":                          30024,
		"BillingPurchaseFreeProduct":                     30025,
		"BillingProductSelectionSlotEmpty":               30026,
		"BillingProductSelectionSlotNotMatch":            30027,
		"BillingProductSelectConditionFailed":            30028,
		"BillingProductSelectionSlotNotFound":            30029,
		"ClanNotFound":                                   31000,
		"ClanSearchFailed":                               31001,
		"ClanEmptySearchString":                          31002,
		"ClanAccountAlreadyJoinedClan":                   31003,
		"ClanAccountAlreadyQuitClan":                     31004,
		"ClanCreateFailed":                               31005,
		"ClanMemberExceedCapacity":                       31006,
		"ClanDoesNotHavePermission":                      31007,
		"ClanTargetAccountIsNotApplicant":                31008,
		"ClanMemberNotFound":                             31009,
		"ClanCanNotKick":                                 31010,
		"ClanCanNotDismiss":                              31011,
		"ClanCanNotQuit":                                 31012,
		"ClanRejoinCoolOff":                              31013,
		"ClanChangeMemberGradeFailed":                    31014,
		"ClanHasBeenDisMissed":                           31015,
		"ClanCannotChangeJoinOption":                     31016,
		"ClanExceedConferCountLimit":                     31017,
		"ClanBusy":                                       31018,
		"ClanNameEmptyString":                            31019,
		"ClanNameWithInvalidLength":                      31020,
		"ClanAssistCharacterAlreadyDeployed":             31021,
		"ClanAssistNotValidUse":                          31022,
		"ClanAssistCharacterChanged":                     31023,
		"ClanAssistCoolTime":                             31024,
		"ClanAssistAlreadyUsedInRaidRoom":                31025,
		"ClanAssistAlreadyUsedInTimeAttackDungeonRoom":   31026,
		"ClanAssistEchelonHasAssistOnly":                 31027,
		"PaymentInvalidSign":                             32000,
		"PaymentInvalidSeed1":                            32001,
		"PaymentInvalidSeed2":                            32002,
		"PaymentInvalidInput":                            32003,
		"PaymentNotFoundPurchase":                        32004,
		"PaymentGetPurchaseOrderNotZero":                 32005,
		"PaymentSetPurchaseOrderNotZero":                 32006,
		"PaymentException":                               32007,
		"PaymentInvalidState":                            32008,
		"SessionNotFound":                                33000,
		"SessionParseFail":                               33001,
		"SessionInvalidInput":                            33002,
		"SessionNotAuth":                                 33003,
		"SessionDuplicateLogin":                          33004,
		"SessionTimeOver":                                33005,
		"SessionInvalidVersion":                          33006,
		"SessionChangeDate":                              33007,
		"CallName_RenameCoolTime":                        34000,
		"CallName_EmptyString":                           34001,
		"CallName_InvalidString":                         34002,
		"CallName_TTSServerIsNotAvailable":               34003,
		"CouchbaseInvalidCas":                            35000,
		"CouchbaseOperationFailed":                       35001,
		"CouchbaseRollBackFailed":                        35002,
		"EventContentCannotSelectBuff":                   36000,
		"EventContentNoBuffGroupAvailable":               36001,
		"EventContentBuffGroupIdDuplicated":              36002,
		"EventContentNotOpen":                            36003,
		"EventContentNoTotalRewardAvailable":             36004,
		"EventContentBoxGachaPurchaseFailed":             36005,
		"EventContentBoxGachaCannotRefresh":              36006,
		"EventContentCardShopCannotShuffle":              36007,
		"EventContentElementDoesNotExist":                36008,
		"EventContentElementAlreadyPurchased":            36009,
		"EventContentLocationNotFound":                   36010,
		"EventContentLocationScheduleCanNotAttend":       36011,
		"EventContentDiceRaceDataNotFound":               36012,
		"EventContentDiceRaceAlreadyReceiveLapRewardAll": 36013,
		"EventContentDiceRaceInvalidDiceRaceResultType":  36014,
		"EventContentTreasureDataNotFound":               36015,
		"EventContentTreasureNotComplete":                36016,
		"EventContentTreasureFlipFailed":                 36017,
		"MiniGameStageIsNotOpen":                         37000,
		"MiniGameStageInvalidResult":                     37001,
		"MiniGameShootingStageInvlid":                    37002,
		"MiniGameShootingCannotSweep":                    37003,
		"MiniGameTableBoardSaveNotExist":                 37004,
		"MiniGameTableBoardPlayerCannotMove":             37005,
		"MiniGameTableBoardNoActiveEncounter":            37006,
		"MiniGameTableBoardInvalidEncounterRequest":      37007,
		"MiniGameTableBoardProcessEncounterFailed":       37008,
		"MiniGameTableBoardItemNotExist":                 37009,
		"MiniGameTableBoardInvalidItemUse":               37010,
		"MiniGameTableBoardInvalidClearThemaRequest":     37011,
		"MiniGameTableBoardInvalidSeason":                37012,
		"MiniGameTableBoardInvalidResurrectRequest":      37013,
		"MiniGameTableBoardSweepConditionFail":           37014,
		"MiniGameTableBoardInvalidData":                  37015,
		"MiniGameDreamCannotStartNewGame":                37016,
		"MiniGameDreamCannotApplyMultiplier":             37017,
		"MiniGameDreamCannotReset":                       37018,
		"MiniGameDreamNotEnoughActionCount":              37019,
		"MiniGameDreamSaveNotExist":                      37020,
		"MiniGameDreamActionCountRemain":                 37021,
		"MiniGameDreamRoundNotComplete":                  37022,
		"MiniGameDreamRewardAlreadyReceived":             37023,
		"MiniGameDreamRoundCompleted":                    37024,
		"MiniGameShouldReceiveEndingReward":              37025,
		"MiniGameDefenseCannotUseCharacter":              37026,
		"MiniGameDefenseNotOpenStage":                    37027,
		"MiniGameDefenseCannotApplyMultiplier":           37028,
		"MiniGameRoadPuzzleInvalidTilePlacement":         37029,
		"MiniGameRoadPuzzleCannotTrainDeparture":         37030,
		"MiniGameRoadPuzzleAlreadyCleared":               37031,
		"MiniGameRoadPuzzleCannotSave":                   37032,
		"MiniGameCCGPlayingSaveAlreadyExists":            37033,
		"MiniGameCCGSaveNotExists":                       37034,
		"MiniGameCCGPlayingStageAlreadyExists":           37035,
		"MiniGameCCGPlayingStageNotExists":               37036,
		"MiniGameCCGInvalidOperation":                    37037,
		"MiniGameCCGSaveNotComplete":                     37038,
		"MiniGameCCGNoRerollPoint":                       37039,
		"ProofTokenNotSubmitted":                         38000,
		"SchoolDungeonInfoNotFound":                      39000,
		"SchoolDungeonNotOpened":                         39001,
		"SchoolDungeonInvalidSaveData":                   39002,
		"SchoolDungeonBattleWinnerInvalid":               39003,
		"SchoolDungeonInvalidReward":                     39004,
		"TimeAttackDungeonDataNotFound":                  40000,
		"TimeAttackDungeonNotOpen":                       40001,
		"TimeAttackDungeonRoomTimeOut":                   40002,
		"TimeAttackDungeonRoomPlayCountOver":             40003,
		"TimeAttackDungeonRoomAlreadyExists":             40004,
		"TimeAttackDungeonRoomAlreadyClosed":             40005,
		"TimeAttackDungeonRoomNotExist":                  40006,
		"TimeAttackDungeonInvalidRequest":                40007,
		"TimeAttackDungeonInvalidData":                   40008,
		"WorldRaidDataNotFound":                          41000,
		"WorldRaidSeasonNotOpen":                         41001,
		"WorldRaidBossGroupNotOpen":                      41002,
		"WorldRaidInvalidOpenCondition":                  41003,
		"WorldRaidDifficultyNotOpen":                     41004,
		"WorldRaidAssistCharacterLimitOver":              41005,
		"WorldRaidContainBlackListCharacter":             41006,
		"WorldRaidValidFixedEchelonSetting":              41007,
		"WorldRaidAlredayReceiveRewardAll":               41008,
		"WorldRaidCannotReceiveReward":                   41009,
		"WorldRaidBossAlreadyDead":                       41010,
		"WorldRaidNotAnotherBossKilled":                  41011,
		"WorldRaidBattleResultUpdateFailed":              41012,
		"WorldRaidGemEnterCountLimitOver":                41013,
		"WorldRaidCannotGemEnter":                        41014,
		"WorldRaidNeedClearScenarioBoss":                 41015,
		"WorldRaidBossIsAlive":                           41016,
		"ConquestDataNotFound":                           42000,
		"ConquestAlreadyConquested":                      42001,
		"ConquestNotFullyConquested":                     42002,
		"ConquestStepNotOpened":                          42003,
		"ConquestUnableToReach":                          42004,
		"ConquestUnableToAttack":                         42005,
		"ConquestEchelonChangedCountMax":                 42006,
		"ConquestEchelonNotFound":                        42007,
		"ConquestCharacterAlreadyDeployed":               42008,
		"ConquestMaxUpgrade":                             42009,
		"ConquestUnitNotFound":                           42010,
		"ConquestObjectNotFound":                         42011,
		"ConquestCalculateRewardNotFound":                42012,
		"ConquestInvalidTileType":                        42013,
		"ConquestInvalidObjectType":                      42014,
		"ConquestInvalidSaveData":                        42015,
		"ConquestMaxAssistCountReached":                  42016,
		"ConquestErosionConditionNotSatisfied":           42017,
		"ConquestAdditionalContentNotInUse":              42018,
		"ConquestCannotUseManageEchelon":                 42019,
		"FriendUserIsNotFriend":                          43000,
		"FriendFailedToCreateFriendIdCard":               43001,
		"FriendRequestNotFound":                          43002,
		"FriendInvalidFriendCode":                        43003,
		"FriendAlreadyFriend":                            43004,
		"FriendMaxSentRequestReached":                    43005,
		"FriendMaxReceivedRequestReached":                43006,
		"FriendCannotRequestMaxFriendCountReached":       43007,
		"FriendCannotAcceptMaxFriendCountReached":        43008,
		"FriendOpponentMaxFriendCountReached":            43009,
		"FriendTargetIsBusy":                             43010,
		"FriendRequestTargetIsYourself":                  43011,
		"FriendSearchTargetIsYourself":                   43012,
		"FriendInvalidBackgroundId":                      43013,
		"FriendIdCardCommentLengthOverLimit":             43014,
		"FriendBackgroundNotOwned":                       43015,
		"FriendBlockTargetIsYourself":                    43016,
		"FriendBlockTargetIsAlreadyBlocked":              43017,
		"FriendBlockTargetIsExceedMaxCount":              43018,
		"FriendBlockUserCannotOpenProfile":               43019,
		"FriendBlockUserCannotSendRequest":               43020,
		"EliminateStageIsNotOpened":                      44000,
		"MultiSweepPresetDocumentNotFound":               45000,
		"MultiSweepPresetNameEmpty":                      45001,
		"MultiSweepPresetInvalidStageId":                 45002,
		"MultiSweepPresetInvalidId":                      45003,
		"MultiSweepPresetNameInvalidLength":              45004,
		"MultiSweepPresetTooManySelectStageId":           45005,
		"MultiSweepPresetInvalidSweepCount":              45006,
		"MultiSweepPresetTooManySelectParcelId":          45007,
		"EmblemDataNotFound":                             46000,
		"EmblemAttachFailed":                             46001,
		"EmblemCannotReceive":                            46002,
		"EmblemPassCheckEmblemIsEmpty":                   46003,
		"StickerDataNotFound":                            47000,
		"StickerNotAcquired":                             47001,
		"StickerDocumentNotFound":                        47002,
		"StickerAlreadyUsed":                             47003,
		"ClearDeckInvalidKey":                            48000,
		"ClearDeckOutOfDate":                             48001,
		"FieldDataNotFound":                              60000,
		"FieldInteracionFailed":                          60001,
		"FieldQuestClearFailed":                          60002,
		"FieldInvalidSceneChangedRequest":                60003,
		"FieldInvalidEndDateRequest":                     60004,
		"FieldCreateDailyQuestFailed":                    60005,
		"FieldResetReplayFailed":                         60006,
		"FieldIncreaseMasteryFailed":                     60007,
		"FieldStageDataInvalid":                          60008,
		"FieldStageEnterFail":                            60009,
		"FieldContentIsClosed":                           60010,
		"FieldEventStageNotCleared":                      60011,
		"MultiFloorRaidSeasonNotOpened":                  49000,
		"MultiFloorRaidDataNotFound":                     49001,
		"MultiFloorRaidAssistCharacterLimitOver":         49002,
		"MultiFloorRaidStageOpenConditionFail":           49003,
		"MultiFloorRaidInvalidSummary":                   49004,
		"MultiFloorRaidInvalidRewardRequest":             49005,
		"BattlePassSeasonNotOpen":                        50000,
		"BattlePassBuyLevelAlreadyMaxLevel":              50001,
		"BattlePassBuyLevelMaxLevelOver":                 50002,
		"BattlePassBuyLevelBuyCountError":                50003,
		"BattlePassAlreadyGetRewardAll":                  50004,
	}
)

func (x WebAPIErrorCode) String() string {
	return WebAPIErrorCode_name[int32(x)]
}

func (x WebAPIErrorCode) Value(sr string) WebAPIErrorCode {
	return WebAPIErrorCode(WebAPIErrorCode_value[sr])
}
