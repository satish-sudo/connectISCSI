package main

import (
		"fmt"
		iscsi "connector" 
)

func main() {
	targetPortal := "192.168.x.x"
	vmName := "satish-backup-demo-centos"
	iscsiConnector := iscsi.NewISCSIConnector(targetPortal, vmName)

	err := iscsiConnector.CheckForInitiator()
	if err != nil {
		err = iscsiConnector.SetServiceStartupType(); if err!= nil { fmt.Printf("SetServiceStartupType Failed: Err:%s", err) }
		err = iscsiConnector.StartService(); if err!= nil { fmt.Printf("StartService Failed: Err:%s", err) }
	}
	
	err = iscsiConnector.AddTargetPortal()
	if err == nil {
		fmt.Println("AddTargetPortal successfully executed")
	}	
	
	err, targetIqn := iscsiConnector.GetTargetsList()
	if len(targetIqn) == 0 {
		fmt.Println("No target found.")
		//return errors.New("No target found")
	}
	if err == nil {
		fmt.Println("GetTargetsList successfully executed.")
		for index, target := range targetIqn {
				fmt.Printf("Connecting to target[%d]: %s\n",index, target)
			err = iscsiConnector.ConnectToTarget(target)
			if err == nil {
				fmt.Printf("Successfully connected to target %s\n",target)
			}
		}
	}
}	
	