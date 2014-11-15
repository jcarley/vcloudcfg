package datastore_test

import (
	"github.com/jcarley/vcloudcfg/datastore"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Datastore", func() {

	Describe("Load function", func() {

		It("loads the Boxfile.json file", func() {
			ds := datastore.NewDatastore()
			err := ds.Load()
			Expect(err).NotTo(HaveOccurred())
			// fmt.Printf("%v-", ds.Boxes)
		})

	})

})
