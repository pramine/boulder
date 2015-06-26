// Copyright 2015 ISRG.  All rights reserved
// This Source Code Form is subject to the terms of the Mozilla Public
// License, v. 2.0. If a copy of the MPL was not distributed with this
// file, You can obtain one at http://mozilla.org/MPL/2.0/.

package test

import (
	"database/sql"
	"time"
	"net"

	// Load SQLite3 for test purposes
	_ "github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/mattn/go-sqlite3"
	gorp "github.com/letsencrypt/boulder/Godeps/_workspace/src/gopkg.in/gorp.v1"
	"github.com/letsencrypt/boulder/Godeps/_workspace/src/github.com/miekg/dns"
)

// MockCADatabase is a mock
type MockCADatabase struct {
	db    *gorp.DbMap
	count int64
}

// NewMockCertificateAuthorityDatabase is a mock
func NewMockCertificateAuthorityDatabase() (mock *MockCADatabase, err error) {
	db, err := sql.Open("sqlite3", ":memory:")
	dbmap := &gorp.DbMap{Db: db, Dialect: gorp.SqliteDialect{}}
	mock = &MockCADatabase{db: dbmap, count: 1}
	return mock, err
}

// Begin is a mock
func (cadb *MockCADatabase) Begin() (*gorp.Transaction, error) {
	return cadb.db.Begin()
}

// IncrementAndGetSerial is a mock
func (cadb *MockCADatabase) IncrementAndGetSerial(*gorp.Transaction) (int64, error) {
	cadb.count = cadb.count + 1
	return cadb.count, nil
}

// CreateTablesIfNotExists is a mock
func (cadb *MockCADatabase) CreateTablesIfNotExists() error {
	return nil
}

// MockDNS is a mock
type MockDNS struct {
}

// ExchangeOne is a mock
func (mock *MockDNS) ExchangeOne(m *dns.Msg) (rsp *dns.Msg, rtt time.Duration, err error) {
	return m, 0, nil
}

// LookupTXT is a mock
func (mock *MockDNS) LookupTXT(hostname string) ([]string, time.Duration, error) {
	return []string{"hostname"}, 0, nil
}

// LookupDNSSEC is a mock
func (mock *MockDNS) LookupDNSSEC(m *dns.Msg) (*dns.Msg, time.Duration, error) {
	return m, 0, nil
}

// LookupHost is a mock
func (mock *MockDNS) LookupHost(hostname string) ([]net.IP, time.Duration, error) {
	return nil, 0, nil
}

// LookupCNAME is a mock
func (mock *MockDNS) LookupCNAME(domain string) (string, error) {
	return "hostname", nil
}

// LookupCAA is a mock
func (mock *MockDNS) LookupCAA(domain string, alias bool) ([]*dns.CAA, error) {
	return []*dns.CAA{}, nil
}
