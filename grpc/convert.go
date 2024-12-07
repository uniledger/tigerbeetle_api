package grpc

import (
	"github.com/lil5/tigerbeetle_api/proto/tigerbeetle"
	"github.com/lil5/tigerbeetle_api/shared"
	"github.com/samber/lo"
	"github.com/tigerbeetle/tigerbeetle-go/pkg/types"
)

func AccountToProtoAccount(tbAccount types.Account) *proto.Account {
    tbFlags := tbAccount.AccountFlags()
    pFlags := proto.AccountFlags{
        Linked:                     lo.ToPtr(tbFlags.Linked),
        DebitsMustNotExceedCredits: lo.ToPtr(tbFlags.DebitsMustNotExceedCredits),
        CreditsMustNotExceedDebits: lo.ToPtr(tbFlags.CreditsMustNotExceedDebits),
        History:                    lo.ToPtr(tbFlags.History),
    }
    
    return &proto.Account{
        Id:             tbAccount.ID.String(),
        DebitsPending:  tbAccount.DebitsPending.String(),
        DebitsPosted:   tbAccount.DebitsPosted.String(),
        CreditsPending: tbAccount.CreditsPending.String(),
        CreditsPosted:  tbAccount.CreditsPosted.String(),
        UserData128:    tbAccount.UserData128.String(),
        UserData64:     int64(tbAccount.UserData64),
        UserData32:     int32(tbAccount.UserData32),
        Ledger:         int64(tbAccount.Ledger),
        Code:           int32(tbAccount.Code),
        Flags:          &pFlags,
        Timestamp:      shared.TimestampFromUintToString(tbAccount.Timestamp),
    }
}


func TransferToProtoTransfer(tbTransfer types.Transfer) *proto.Transfer {
    tbFlags := tbTransfer.TransferFlags()
    pFlags := &proto.TransferFlags{
        Linked:              lo.ToPtr(tbFlags.Linked),
        Pending:             lo.ToPtr(tbFlags.Pending),
        PostPendingTransfer: lo.ToPtr(tbFlags.PostPendingTransfer),
        VoidPendingTransfer: lo.ToPtr(tbFlags.VoidPendingTransfer),
        BalancingDebit:      lo.ToPtr(tbFlags.BalancingDebit),
        BalancingCredit:     lo.ToPtr(tbFlags.BalancingCredit),
    }

    var pendingId *string
    emptyUint128 := types.Uint128{}
    if tbTransfer.PendingID != emptyUint128 {
        id := tbTransfer.PendingID.String()
        pendingId = &id
    }

    return &proto.Transfer{
        Id:              tbTransfer.ID.String(),
        DebitAccountId:  tbTransfer.DebitAccountID.String(),
        CreditAccountId: tbTransfer.CreditAccountID.String(),
        Amount:          tbTransfer.Amount.String(),
        PendingId:       pendingId,
        UserData128:     tbTransfer.UserData128.String(),
        UserData64:      int64(tbTransfer.UserData64),
        UserData32:      int32(tbTransfer.UserData32),
        Ledger:         int64(tbTransfer.Ledger),
        Code:           int32(tbTransfer.Code),
        TransferFlags:   pFlags,
        Timestamp:       lo.ToPtr(shared.TimestampFromUintToString(tbTransfer.Timestamp)),
    }
}

func AccountFilterFromProtoToTigerbeetle(pAccountFilter *proto.AccountFilter) (*types.AccountFilter, error) {
	accountID, err := shared.HexStringToUint128(pAccountFilter.AccountId)
	if err != nil {
		return nil, err
	}

	timestampMin, err := shared.TimestampFromPstringToUint(pAccountFilter.TimestampMin)
	if err != nil {
		return nil, err
	}
	timestampMax, err := shared.TimestampFromPstringToUint(pAccountFilter.TimestampMax)
	if err != nil {
		return nil, err
	}

	var tbFlags types.AccountFilterFlags
	if pAccountFilter.Flags != nil {
		tbFlags = types.AccountFilterFlags{
			Debits:   lo.FromPtrOr(pAccountFilter.Flags.Debits, false),
			Credits:  lo.FromPtrOr(pAccountFilter.Flags.Credits, false),
			Reversed: lo.FromPtrOr(pAccountFilter.Flags.Reserved, false),
		}
	}

	return &types.AccountFilter{
		AccountID:    *accountID,
		TimestampMin: *timestampMin,
		TimestampMax: *timestampMax,
		Limit:        uint32(pAccountFilter.Limit),
		Flags:        tbFlags.ToUint32(),
	}, nil
}

func AccountBalanceFromTigerbeetleToProto(tbBalance types.AccountBalance) *proto.AccountBalance {
    return &proto.AccountBalance{
        DebitsPending:  tbBalance.DebitsPending.String(),
        DebitsPosted:   tbBalance.DebitsPosted.String(),
        CreditsPending: tbBalance.CreditsPending.String(),
        CreditsPosted:  tbBalance.CreditsPosted.String(),
        Timestamp:      shared.TimestampFromUintToString(tbBalance.Timestamp),
    }
}