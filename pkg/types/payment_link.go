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

import (
	"encoding/json"
	"errors"
)

type PaymentLink struct {
	ID        string                `json:"id"`
	Amount    int                   `json:"amount"`
	Checkouts []PaymentLinkCheckout `json:"checkouts"`
	Closed    bool                  `json:"closed"`
	Code      string                `json:"code"`
	Currency  string                `json:"currency"`
	Customer  PaymentLinkCustomer   `json:"customer"`
	Items     []PaymentLinkItem     `json:"items"`
	SessionID string                `json:"session_id"`
	Status    string                `json:"status"`
	CreatedAt string                `json:"created_at"`
	UpdatedAt string                `json:"updated_at"`
}

type PaymentLinkCheckout struct {
	ID                          string                `json:"id"`
	AcceptedMultiPaymentMethods []string              `json:"accepted_multi_payment_methods"`
	AcceptedPaymentMethods      []string              `json:"accepted_payment_methods"`
	Amount                      int                   `json:"amount"`
	BillingAddress              json.RawMessage       `json:"billing_address"`
	BillingAddressEditable      bool                  `json:"billing_address_editable"`
	CreditCard                  PaymentLinkCreditCard `json:"credit_card"`
	Currency                    string                `json:"currency"`
	Customer                    PaymentLinkCustomer   `json:"customer"`
	CustomerEditable            bool                  `json:"customer_editable"`
	ExpiresAt                   string                `json:"expires_at"`
	Metadata                    json.RawMessage       `json:"metadata"`
	PaymentURL                  string                `json:"payment_url"`
	RequiredFields              []string              `json:"required_fields"`
	Shippable                   bool                  `json:"shippable"`
	SkipCheckoutSuccessPage     bool                  `json:"skip_checkout_success_page"`
	Status                      string                `json:"status"`
	SuccessURL                  string                `json:"success_url"`
	CreatedAt                   string                `json:"created_at"`
	UpdatedAt                   string                `json:"updated_at"`
}

type PaymentLinkCreditCard struct {
	Authentication PaymentLinkCreditCardAuth          `json:"authentication"`
	Capture        bool                               `json:"capture"`
	Installments   []PaymentLinkCreditCardInstallment `json:"installments"`
}

type PaymentLinkCreditCardAuth struct {
	ThreedSecure json.RawMessage `json:"threed_secure"`
	Type         string          `json:"type"`
}

type PaymentLinkCreditCardInstallment struct {
	Number int `json:"number"`
	Total  int `json:"total"`
}

type PaymentLinkCustomer struct {
	ID         string          `json:"id"`
	Name       string          `json:"name"`
	Delinquent bool            `json:"delinquent"`
	Phones     json.RawMessage `json:"phones"`
	CreatedAt  string          `json:"created_at"`
	UpdatedAt  string          `json:"updated_at"`
}

type PaymentLinkItem struct {
	ID          string `json:"id"`
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	Status      string `json:"status"`
	Type        string `json:"type"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type PaymentLinkInput struct {
	AccountID string                    `json:"account_id"`
	Items     []PaymentLinkItemInput    `json:"items"`
	Customer  PaymentLinkCustomerInput  `json:"customer"`
	Closed    bool                      `json:"closed"`
	Payments  []PaymentLinkPaymentInput `json:"payments"`
}

type PaymentLinkItemInput struct {
	Amount      int    `json:"amount"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
}

type PaymentLinkCustomerInput struct {
	Name string `json:"name"`
}

type PaymentLinkPaymentInput struct {
	Amount        int                      `json:"amount"`
	PaymentMethod string                   `json:"payment_method"`
	Checkout      PaymentLinkCheckoutInput `json:"checkout"`
}

type PaymentLinkCheckoutInput struct {
	ExpiresIn                   string                     `json:"expires_in"`
	SkipCheckoutSuccessPage     bool                       `json:"skip_checkout_success_page"`
	BillingAddressEditable      bool                       `json:"billing_address_editable"`
	CustomerEditable            bool                       `json:"customer_editable"`
	AcceptedPaymentMethods      []string                   `json:"accepted_payment_methods"`
	AcceptedMultiPaymentMethods []string                   `json:"accepted_multi_payment_methods"`
	SuccessURL                  string                     `json:"success_url"`
	CreditCard                  PaymentLinkCreditCardInput `json:"credit_card"`
}

type PaymentLinkCreditCardInput struct {
	Capture      bool                               `json:"capture"`
	Installments []PaymentLinkCreditCardInstallment `json:"installments"`
}

type PaymentLinkCancelInput struct {
	AccountID string `json:"account_id"`
	Status    string `json:"status"`
}

func (p PaymentLinkCancelInput) Validate() error {
	if p.AccountID == "" {
		return errors.New("account_id can't be empty")
	}

	if p.Status != "canceled" {
		return errors.New("status must be equal to canceled")
	}

	return nil
}
