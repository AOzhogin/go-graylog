package testdata

import (
	"github.com/suzuki-shunsuke/go-set"

	"github.com/suzuki-shunsuke/go-graylog/v8"
)

var (
	Users = &graylog.UsersBody{
		Users: []graylog.User{
			{
				Username:         "admin",
				Email:            "",
				FullName:         "Administrator",
				Password:         "",
				ID:               "local:admin",
				Timezone:         "UTC",
				LastActivity:     "2019-09-21T01:28:20.552+0000",
				ClientAddress:    "172.18.0.1",
				SessionTimeoutMs: 28800000,
				External:         false,
				ReadOnly:         true,
				SessionActive:    true,
				Preferences: &graylog.Preferences{
					UpdateUnfocussed:  false,
					EnableSmartSearch: true,
				},
				Startpage: nil,
				Roles: set.StrSet{
					"Admin": struct{}{},
				},
				Permissions: set.StrSet{
					"*": struct{}{},
				},
			},
			{
				Username:         "graylog-sidecar",
				Email:            "sidecar@graylog.local",
				FullName:         "Sidecar System User (built-in)",
				Password:         "",
				ID:               "5d84bfbb2ab79c000d35d402",
				Timezone:         "UTC",
				LastActivity:     "",
				ClientAddress:    "",
				SessionTimeoutMs: 28800000,
				External:         false,
				ReadOnly:         false,
				SessionActive:    false,
				Preferences: &graylog.Preferences{
					UpdateUnfocussed:  false,
					EnableSmartSearch: true,
				},
				Startpage: nil,
				Roles: set.StrSet{
					"Reader":                    struct{}{},
					"Sidecar System (Internal)": struct{}{},
				},
				Permissions: set.StrSet{
					"buffers:read":                          struct{}{},
					"clusterconfigentry:read":               struct{}{},
					"decorators:read":                       struct{}{},
					"fieldnames:read":                       struct{}{},
					"indexercluster:read":                   struct{}{},
					"inputs:read":                           struct{}{},
					"journal:read":                          struct{}{},
					"jvmstats:read":                         struct{}{},
					"messagecount:read":                     struct{}{},
					"messages:analyze":                      struct{}{},
					"messages:read":                         struct{}{},
					"metrics:read":                          struct{}{},
					"savedsearches:create":                  struct{}{},
					"savedsearches:edit":                    struct{}{},
					"savedsearches:read":                    struct{}{},
					"sidecar_collector_configurations:read": struct{}{},
					"sidecar_collectors:read":               struct{}{},
					"sidecars:update":                       struct{}{},
					"system:read":                           struct{}{},
					"throughput:read":                       struct{}{},
					"users:edit:graylog-sidecar":            struct{}{},
					"users:passwordchange:graylog-sidecar":  struct{}{},
					"users:tokencreate:graylog-sidecar":     struct{}{},
					"users:tokenlist:graylog-sidecar":       struct{}{},
					"users:tokenremove:graylog-sidecar":     struct{}{},
				},
			},
			{
				Username:         "test",
				Email:            "test@example.com",
				FullName:         "test test",
				Password:         "",
				ID:               "5d84c1a92ab79c000d35d6cb",
				Timezone:         "",
				LastActivity:     "",
				ClientAddress:    "",
				SessionTimeoutMs: 28800000,
				External:         false,
				ReadOnly:         false,
				SessionActive:    false,
				Preferences: &graylog.Preferences{
					UpdateUnfocussed:  false,
					EnableSmartSearch: true,
				},
				Startpage: nil,
				Roles: set.StrSet{
					"Reader": struct{}{},
				},
				Permissions: set.StrSet{
					"buffers:read":              struct{}{},
					"clusterconfigentry:read":   struct{}{},
					"decorators:read":           struct{}{},
					"fieldnames:read":           struct{}{},
					"indexercluster:read":       struct{}{},
					"inputs:read":               struct{}{},
					"journal:read":              struct{}{},
					"jvmstats:read":             struct{}{},
					"messagecount:read":         struct{}{},
					"messages:analyze":          struct{}{},
					"messages:read":             struct{}{},
					"metrics:read":              struct{}{},
					"savedsearches:create":      struct{}{},
					"savedsearches:edit":        struct{}{},
					"savedsearches:read":        struct{}{},
					"system:read":               struct{}{},
					"throughput:read":           struct{}{},
					"users:edit:test":           struct{}{},
					"users:passwordchange:test": struct{}{},
					"users:tokencreate:test":    struct{}{},
					"users:tokenlist:test":      struct{}{},
					"users:tokenremove:test":    struct{}{},
				},
			},
		},
	}
)
