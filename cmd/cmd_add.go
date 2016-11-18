/**
 * clinot.es client
 * Copyright (C) 2016 Sebastian Müller
 *
 * This program is free software: you can redistribute it and/or modify
 * it under the terms of the GNU General Public License as published by
 * the Free Software Foundation, either version 3 of the License, or
 * (at your option) any later version.
 *
 * This program is distributed in the hope that it will be useful,
 * but WITHOUT ANY WARRANTY; without even the implied warranty of
 * MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
 * GNU General Public License for more details.
 *
 * You should have received a copy of the GNU General Public License
 * along with this program. If not, see <http://www.gnu.org/licenses/>.
 */

package cmd

import (
	"strings"

	fb "github.com/sbstjn/feedback"
	"github.com/spf13/cobra"
)

type jsonDataAdd struct {
	Address string
	Token   string
	Note    string
}

func addHandler(cmd *cobra.Command, args []string) {
	// Fail when APIAddress or APIToken is missing
	ensureCredentials()

	note := strings.Join(args, " ")
	if len(note) > 100 {
		fb.Fail("Note must not be longer than 100 characters!")
	}

	data := jsonDataAdd{APIAddress, APIToken, note}
	if err := newRequest("/add").post(data); err != nil {
		fb.Fail("Fail")
	}

	fb.Done("Done")
}

func init() {
	RootCmd.AddCommand(&cobra.Command{
		Use:   "add",
		Short: "Add a note",
		Run:   addHandler,
	})
}
