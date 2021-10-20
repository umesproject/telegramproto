// Copyright (c) 2020-2021 KHS Films
//
// This file is a part of mtproto package.
// See https://github.com/umesproject/mtproto/blob/master/LICENSE for details

package objects

import (
	"github.com/umesproject/mtproto/internal/encoding/tl"
)

func init() {
	tl.RegisterObjects(
		&ReqPQParams{},
		&ReqDHParamsParams{},
		&SetClientDHParamsParams{},
		&PingParams{},
		&ResPQ{},
		&PQInnerData{},
		&ServerDHParamsFail{},
		&ServerDHParamsOk{},
		&ServerDHInnerData{},
		&ClientDHInnerData{},
		&DHGenOk{},
		&DHGenRetry{},
		&DHGenFail{},
		&RpcResult{},
		&RpcError{},
		&RpcAnswerUnknown{},
		&RpcAnswerDroppedRunning{},
		&RpcAnswerDropped{},
		&FutureSaltObj{},
		&FutureSaltsObj{},
		&Pong{},
		&NewSessionCreated{},
		&MessageContainer{},
		&MsgCopy{},
		&GzipPacked{},
		&MsgsAck{},
		&BadMsgNotification{},
		&BadServerSalt{},
		&MsgResendReq{},
		&MsgsStateReq{},
		&MsgsStateInfo{},
		&MsgsAllInfo{},
		&MsgsDetailedInfo{},
		&MsgsNewDetailedInfo{},
	)
}