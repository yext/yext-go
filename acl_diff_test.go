package yext_test

// func TestACL_Diff(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		aclA      yext.ACL
// 		aclB      yext.ACL
// 		wantDelta *yext.ACL
// 		wantDiff  bool
// 	}{
// 		{
// 			name: "Identical ACLs",
// 			aclA: yext.ACL{
// 				Role: yext.Role{
// 					Id:   yext.String("3"),
// 					Name: yext.String("Example Role"),
// 				},
// 				On:       "12345",
// 				AccessOn: yext.ACCESS_FOLDER,
// 			},
// 			aclB: yext.ACL{
// 				Role: yext.Role{
// 					Id:   yext.String("3"),
// 					Name: yext.String("Example Role"),
// 				},
// 				On:       "12345",
// 				AccessOn: yext.ACCESS_FOLDER,
// 			},
// 			wantDelta: nil,
// 			wantDiff:  false,
// 		},
// 		{
// 			name: "Different Roles in ACL",
// 			aclA: yext.ACL{
// 				Role: yext.Role{
// 					Id:   yext.String("3"),
// 					Name: yext.String("Example Role"),
// 				},
// 				On:       "12345",
// 				AccessOn: yext.ACCESS_FOLDER,
// 			},
// 			aclB: yext.ACL{
// 				Role: yext.Role{
// 					Id:   yext.String("4"),
// 					Name: yext.String("Example Role Two"),
// 				},
// 				On:       "12345",
// 				AccessOn: yext.ACCESS_FOLDER,
// 			},
// 			wantDelta: &yext.ACL{
// 				Role: yext.Role{
// 					Id:   yext.String("4"),
// 					Name: yext.String("Example Role Two"),
// 				},
// 			},
// 			wantDiff: true,
// 		},
// 		{
// 			name: "Different 'On' params in ACL",
// 			aclA: yext.ACL{
// 				Role: yext.Role{
// 					Id:   yext.String("3"),
// 					Name: yext.String("Example Role"),
// 				},
// 				On:       "12345",
// 				AccessOn: yext.ACCESS_FOLDER,
// 			},
// 			aclB: yext.ACL{
// 				Role: yext.Role{
// 					Id:   yext.String("3"),
// 					Name: yext.String("Example Role"),
// 				},
// 				On:       "123456",
// 				AccessOn: yext.ACCESS_FOLDER,
// 			},
// 			wantDelta: &yext.ACL{
// 				On: "123456",
// 			},
// 			wantDiff: true,
// 		},
// 		{
// 			name: "Different 'AccessOn' params in ACL",
// 			aclA: yext.ACL{
// 				Role: yext.Role{
// 					Id:   yext.String("3"),
// 					Name: yext.String("Example Role"),
// 				},
// 				On:       "12345",
// 				AccessOn: yext.ACCESS_FOLDER,
// 			},
// 			aclB: yext.ACL{
// 				Role: yext.Role{
// 					Id:   yext.String("3"),
// 					Name: yext.String("Example Role"),
// 				},
// 				On:       "12345",
// 				AccessOn: yext.ACCESS_LOCATION,
// 			},
// 			wantDelta: &yext.ACL{
// 				AccessOn: yext.ACCESS_LOCATION,
// 			},
// 			wantDiff: true,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		if gotDelta, gotDiff := test.aclA.Diff(test.aclB); !reflect.DeepEqual(test.wantDelta, gotDelta) || test.wantDiff != gotDiff {
// 			t.Error(fmt.Sprintf("test '%s' failed, got diff: %t, wanted diff: %t, got delta: %+v, wanted delta: %+v", test.name, test.wantDiff, gotDiff, test.wantDelta, gotDelta))
// 		}
// 	}
// }
//
// func TestACLList_Diff(t *testing.T) {
// 	tests := []struct {
// 		name      string
// 		aclListA  yext.ACLList
// 		aclListB  yext.ACLList
// 		wantDelta yext.ACLList
// 		wantDiff  bool
// 	}{
// 		{
// 			name: "Identical ACLLists",
// 			aclListA: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("4"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			aclListB: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("4"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			wantDelta: nil,
// 			wantDiff:  false,
// 		},
// 		{
// 			name: "Identical ACLs in ACLLists",
// 			aclListA: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 			},
// 			aclListB: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 			},
// 			wantDelta: nil,
// 			wantDiff:  false,
// 		},
// 		{
// 			name: "Different Length in ACLLists",
// 			aclListA: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("4"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			aclListB: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("4"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("5"),
// 						Name: yext.String("Example Role Three"),
// 					},
// 					On:       "1234567",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			wantDelta: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("4"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("5"),
// 						Name: yext.String("Example Role Three"),
// 					},
// 					On:       "1234567",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			wantDiff: true,
// 		},
// 		{
// 			name: "Different Items in ACLLists",
// 			aclListA: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("4"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			aclListB: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("33"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("44"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			wantDelta: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("33"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("44"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			wantDiff: true,
// 		},
// 		{
// 			name: "Some Identical and Some Different Items in ACLLists",
// 			aclListA: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("4"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			aclListB: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("4"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("5"),
// 						Name: yext.String("Example Role Three"),
// 					},
// 					On:       "1234567",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			wantDelta: yext.ACLList{
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("3"),
// 						Name: yext.String("Example Role"),
// 					},
// 					On:       "12345",
// 					AccessOn: yext.ACCESS_FOLDER,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("4"),
// 						Name: yext.String("Example Role Two"),
// 					},
// 					On:       "123456",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 				yext.ACL{
// 					Role: yext.Role{
// 						Id:   yext.String("5"),
// 						Name: yext.String("Example Role Three"),
// 					},
// 					On:       "1234567",
// 					AccessOn: yext.ACCESS_LOCATION,
// 				},
// 			},
// 			wantDiff: true,
// 		},
// 	}
//
// 	for _, test := range tests {
// 		if gotDelta, gotDiff := test.aclListA.Diff(test.aclListB); !reflect.DeepEqual(test.wantDelta, gotDelta) || test.wantDiff != gotDiff {
// 			t.Error(fmt.Sprintf("test '%s' failed, got diff: %t, wanted diff: %t, got delta: %+v, wanted delta: %+v", test.name, test.wantDiff, gotDiff, test.wantDelta, gotDelta))
// 		}
// 	}
// }
