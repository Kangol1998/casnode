// Copyright 2021 The casbin Authors. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package discuzx

import (
	"fmt"

	"github.com/casbin/casnode/object"
	"github.com/casbin/casnode/util"
)

func addForums() {
	tabs := []*object.Tab{}
	nodes := []*object.Node{}

	forumTree, _ := getForumTree()
	for i, groupForum := range forumTree {
		defaultNode := ""
		if len(groupForum.Forums) != 0 {
			defaultNode = groupForum.Forums[0].Name
		}

		tab := &object.Tab{
			Id:          groupForum.Name,
			Name:        groupForum.Name,
			Sorter:      groupForum.Displayorder,
			Ranking:     groupForum.Fid,
			CreatedTime: util.GetCurrentTime(),
			DefaultNode: defaultNode,
			HomePage:    true,
		}

		tabs = append(tabs, tab)
		fmt.Printf("[%d/%d]: Synced group forum: %s\n", i+1, len(forumTree), groupForum.Name)

		for j, forum := range groupForum.Forums {
			forumNode := &object.Node{
				Id:                forum.Name,
				Name:              forum.Name,
				CreatedTime:       util.GetCurrentTime(),
				Desc:              forum.Name,
				Image:             "https://cdn.v2ex.com/navatar/3b8a/6142/215_xxlarge.png?m=1523190513",
				TabId:             groupForum.Name,
				ParentNode:        "",
				PlaneId:           "",
				Sorter:            forum.Displayorder,
				Ranking:           forum.Fid,
				Hot:               forum.Threads,
				Moderators:        []string{},
				MailingList:       "",
				GoogleGroupCookie: "",
				IsHidden:          forum.Status == 0,
			}
			nodes = append(nodes, forumNode)
			fmt.Printf("\t[%d/%d]: Synced forum: %s\n", j+1, len(groupForum.Forums), forum.Name)

			for k, subForum := range forum.Forums {
				subForumNode := &object.Node{
					Id:                subForum.Name,
					Name:              subForum.Name,
					CreatedTime:       util.GetCurrentTime(),
					Desc:              subForum.Name,
					Image:             "https://cdn.v2ex.com/navatar/3b8a/6142/215_xxlarge.png?m=1523190513",
					TabId:             groupForum.Name,
					ParentNode:        forum.Name,
					PlaneId:           "",
					Sorter:            subForum.Displayorder,
					Ranking:           subForum.Fid,
					Hot:               subForum.Threads,
					Moderators:        []string{},
					MailingList:       "",
					GoogleGroupCookie: "",
					IsHidden:          subForum.Status == 0,
				}
				nodes = append(nodes, subForumNode)
				fmt.Printf("\t\t[%d/%d]: Synced sub forum: %s\n", k+1, len(forum.Forums), subForum.Name)
			}
		}
	}

	defaultNode := ""
	if len(nodes) > 0 {
		defaultNode = nodes[0].Id
	}

	tab := &object.Tab{
		Id:          "all",
		Name:        "全部",
		Sorter:      100,
		CreatedTime: util.GetCurrentTime(),
		DefaultNode: defaultNode,
		HomePage:    true,
	}
	tabs = append(tabs, tab)

	object.AddTabs(tabs)
	object.AddNodes(nodes)
}
