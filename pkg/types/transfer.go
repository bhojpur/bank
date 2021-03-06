package types

// Copyright (c) 2018 Bhojpur Consulting Private Limited, India. All rights reserved.

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

type TransferInput struct {
	AccountID   string  `json:"account_id"`
	Currency    string  `json:"currency"`
	Amount      float64 `json:"amount,omitempty"`
	Description string  `json:"description,omitempty"`
	ScheduledTo string  `json:"scheduled_to,omitempty"`
	Target      Target  `json:"target,omitempty"`
	Type        string
}

type Transfer struct {
	ID                       string  `json:"id,omitempty"`
	Currency                 string  `json:"currency"`
	Amount                   float64 `json:"amount,omitempty"`
	Fee                      int     `json:"fee,omitempty"`
	Target                   Target  `json:"target,omitempty"`
	ApprovedAt               string  `json:"approved_at,omitempty"`
	CreatedAt                string  `json:"created_at,omitempty"`
	RejectedAt               string  `json:"rejected_at,omitempty"`
	FailedAt                 string  `json:"failed_at,omitempty"`
	FailureReasonCode        string  `json:"failure_reason_code,omitempty"`
	FailureReasonDescription string  `json:"failure_reason_description,omitempty"`
	Status                   string  `json:"status,omitempty"`
	Description              string  `json:"description,omitempty"`
	ApprovedBy               string  `json:"approved_by,omitempty"`
	CreatedBy                string  `json:"created_by,omitempty"`
	RejectedBy               string  `json:"rejected_by,omitempty"`
	ApprovalExpiredAt        string  `json:"approval_expired_at,omitempty"`
	CancelledAt              string  `json:"cancelled_at,omitempty"`
	FinishedAt               string  `json:"finished_at,omitempty"`
	ScheduledTo              string  `json:"scheduled_to,omitempty"`

	RefundedAt               string
	RefundReasonCode         string `json:"refund_reason_code,omitempty"`
	RefundReasonDescription  string `json:"refund_reason_description,omitempty"`
	DelayedToNextBusinessDay bool   `json:"delayed_to_next_business_day,omitempty"`
	ScheduledToEffective     string `json:"scheduled_to_effective,omitempty"`
	ScheduledToRequested     string `json:"scheduled_to_requested,omitempty"`
}

type Target struct {
	Account TransferAccount `json:"account"`
	Entity  Entity          `json:"entity"`
}

type TransferAccount struct {
	AccountCode           string `json:"account_code,omitempty"`
	AccountType           string `json:"account_type,omitempty"`
	BranchCode            string `json:"branch_code,omitempty"`
	InstitutionISPB       string `json:"institution_ispb"`
	InstitutionCode       string `json:"institution_code"`
	InstitutionName       string `json:"institution_name,omitempty"`
	InstitutionNumberCode string `json:"institution_number_code"`
}

type Entity struct {
	Name         string `json:"name,omitempty"`
	Document     string `json:"document,omitempty"`
	DocumentType string `json:"document_type,omitempty"`
}
