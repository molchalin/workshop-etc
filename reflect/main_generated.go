package main

import (
	"fmt"
	"strconv"
)

func (v *SportsmanRecord) Scan(record []string, m map[string]int) error {
	{
		idx, ok := m["Name"]
		if !ok {
			return fmt.Errorf("can't find field Name")
		}
		str := record[idx]
		v.Name = str
	}
	{
		idx, ok := m["year"]
		if !ok {
			return fmt.Errorf("can't find field year")
		}
		str := record[idx]
		var err error
		v.Year, err = strconv.Atoi(str)
		if err != nil {
			return err
		}
	}
	{
		idx, ok := m["sport"]
		if !ok {
			return fmt.Errorf("can't find field sport")
		}
		str := record[idx]
		v.Sport = str
	}
	{
		idx, ok := m["earnings"]
		if !ok {
			return fmt.Errorf("can't find field earnings")
		}
		str := record[idx]
		var err error
		v.Earnings, err = strconv.ParseFloat(str, 64)
		if err != nil {
			return err
		}
	}
	return nil
}
