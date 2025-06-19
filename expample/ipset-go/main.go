package main

import (
	"fmt"
	"log"
	"net"

	"github.com/lrh3321/ipset-go"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	hashipType()

	hashnetType()

	hashipv6Type()
}

func hashipType() {
	var setname = "test_hash01"

	// equals `ipset create test_hash02 hash:ip`
	err := ipset.Create(setname, ipset.TypeHashIP, ipset.CreateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = ipset.Destroy(setname)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = ipset.Add(setname, &ipset.Entry{IP: net.IPv4(10, 0, 0, 1).To4()})
	if err != nil {
		log.Fatal(err)
	}

	err = ipset.Add(setname, &ipset.Entry{IP: net.IPv4(10, 0, 0, 5).To4()})
	if err != nil {
		log.Fatal(err)
	}

	set, err := ipset.List(setname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(`Name: %s
Type: %s
Header: family inet hashsize %d maxelem %d
Size in memory: %d
References: %d
Number of entries: %d
Members:
`,
		set.SetName,
		set.TypeName,
		set.HashSize,
		set.MaxElements,
		set.SizeInMemory,
		set.References,
		set.NumEntries,
	)

	for _, e := range set.Entries {
		fmt.Println(e.IP.String())
	}

	/*
	   Name: test_hash01
	   Type: hash:ip
	   Header: family inet hashsize 1024 maxelem 65536
	   Size in memory: 296
	   References: 0
	   Number of entries: 2
	   Members:
	   10.0.0.1
	   10.0.0.5
	*/
}

func hashipv6Type() {
	var setname = "test_hash03"

	// equals `ipset create test_hash03 hash:ip family inet6`
	err := ipset.Create(setname, ipset.TypeHashIP, ipset.CreateOptions{Family: ipset.FamilyIPV6})
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = ipset.Destroy(setname)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = ipset.Add(setname, &ipset.Entry{IP: net.ParseIP("2001:db8:1::1")})
	if err != nil {
		log.Fatal(err)
	}

	err = ipset.Add(setname, &ipset.Entry{IP: net.ParseIP("2001:db7:1::1")})
	if err != nil {
		log.Fatal(err)
	}

	set, err := ipset.List(setname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(`Name: %s
Type: %s
Header: family inet hashsize %d maxelem %d
Size in memory: %d
References: %d
Number of entries: %d
Members:
`,
		set.SetName,
		set.TypeName,
		set.HashSize,
		set.MaxElements,
		set.SizeInMemory,
		set.References,
		set.NumEntries,
	)

	for _, e := range set.Entries {
		fmt.Println(e.IP.String())
	}

	/*
		Name: test_hash03
		Type: hash:ip
		Header: family inet hashsize 1024 maxelem 65536
		Size in memory: 400
		References: 0
		Number of entries: 2
		Members:
		2001:db8:1::1
		2001:db7:1::1
	*/
}

func hashnetType() {
	var setname = "test_hash02"
	err := ipset.Create(setname, ipset.TypeHashNet, ipset.CreateOptions{})
	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		err = ipset.Destroy(setname)
		if err != nil {
			log.Fatal(err)
		}
	}()

	err = ipset.Add(setname, &ipset.Entry{IP: net.IPv4(10, 0, 0, 0).To4(), CIDR: 24})
	if err != nil {
		log.Fatal(err)
	}

	err = ipset.Add(setname, &ipset.Entry{IP: net.IPv4(10, 0, 5, 0).To4(), CIDR: 26})
	if err != nil {
		log.Fatal(err)
	}

	set, err := ipset.List(setname)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf(`Name: %s
Type: %s
Header: family inet hashsize %d maxelem %d
Size in memory: %d
References: %d
Number of entries: %d
Members:
`,
		set.SetName,
		set.TypeName,
		set.HashSize,
		set.MaxElements,
		set.SizeInMemory,
		set.References,
		set.NumEntries,
	)

	for _, e := range set.Entries {
		fmt.Println((&net.IPNet{IP: e.IP, Mask: net.CIDRMask(int(e.CIDR), 32)}).String())
	}

	/*
		Name: test_hash02
		Type: hash:net
		Header: family inet hashsize 1024 maxelem 65536
		Size in memory: 568
		References: 0
		Number of entries: 2
		Members:
		10.0.5.0/26
		10.0.0.0/24
	*/
}
