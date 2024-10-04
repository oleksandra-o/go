package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

// Customer represents a customer at the service station.
type Customer struct {
	ID      int
	Name    string
	Contact string
	Vehicle string
}

// Service represents a service provided by the service station.
type Service struct {
	ID          int
	Name        string
	Description string
	Price       float64
}

// Transaction represents a service transaction.
type Transaction struct {
	ID         int
	CustomerID int
	ServiceID  int
	Date       time.Time
}

// Notification represents a coupon notification.
type Notification struct {
	ID         int
	CustomerID int
	ServiceID  int
	Date       time.Time
}

// ServiceStation represents a service station with its data.
type ServiceStation struct {
	Customers     []Customer
	Services      []Service
	Transactions  []Transaction
	Notifications []Notification
}

func (ss *ServiceStation) LoadData() error {
	// Load customer data from CSV file
	customerFile, err := os.Open("customers.csv")
	if err != nil {
		return err
	}
	defer customerFile.Close()

	customerReader := csv.NewReader(customerFile)
	customerRecords, err := customerReader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range customerRecords {
		id, _ := strconv.Atoi(record[0])
		ss.Customers = append(ss.Customers, Customer{
			ID:      id,
			Name:    record[1],
			Contact: record[2],
			Vehicle: record[3],
		})
	}

	// Load service data from CSV file
	serviceFile, err := os.Open("services.csv")
	if err != nil {
		return err
	}
	defer serviceFile.Close()

	serviceReader := csv.NewReader(serviceFile)
	serviceRecords, err := serviceReader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range serviceRecords {
		id, _ := strconv.Atoi(record[0])
		price, _ := strconv.ParseFloat(record[3], 64)
		ss.Services = append(ss.Services, Service{
			ID:          id,
			Name:        record[1],
			Description: record[2],
			Price:       price,
		})
	}

	// Load transaction data from CSV file
	transactionFile, err := os.Open("transactions.csv")
	if err != nil {
		return err
	}
	defer transactionFile.Close()

	transactionReader := csv.NewReader(transactionFile)
	transactionRecords, err := transactionReader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range transactionRecords {
		id, _ := strconv.Atoi(record[0])
		customerID, _ := strconv.Atoi(record[1])
		serviceID, _ := strconv.Atoi(record[2])
		date, _ := time.Parse("2006-01-02", record[3])
		ss.Transactions = append(ss.Transactions, Transaction{
			ID:         id,
			CustomerID: customerID,
			ServiceID:  serviceID,
			Date:       date,
		})
	}

	// Load notification data from CSV file
	notificationFile, err := os.Open("notifications.csv")
	if err != nil {
		return err
	}
	defer notificationFile.Close()

	notificationReader := csv.NewReader(notificationFile)
	notificationRecords, err := notificationReader.ReadAll()
	if err != nil {
		return err
	}

	for _, record := range notificationRecords {
		id, _ := strconv.Atoi(record[0])
		customerID, _ := strconv.Atoi(record[1])
		serviceID, _ := strconv.Atoi(record[2])
		date, _ := time.Parse("2006-01-02", record[3])
		ss.Notifications = append(ss.Notifications, Notification{
			ID:         id,
			CustomerID: customerID,
			ServiceID:  serviceID,
			Date:       date,
		})
	}

	return nil
}

func (ss *ServiceStation) GenerateServiceReport(date time.Time) {
	fmt.Println("----------------------------------------")
	fmt.Println("Service Report -", date.Format("2006-01-02"))
	fmt.Println("----------------------------------------")
	fmt.Println("Customer Name | Vehicle | Service | Last Date")
	fmt.Println("----------------------------------------")

	for _, transaction := range ss.Transactions {
		if transaction.Date.AddDate(0, 6, 0).Before(date) {
			continue
		}

		customer := ss.getCustomerByID(transaction.CustomerID)
		service := ss.getServiceByID(transaction.ServiceID)

		fmt.Printf("%-14s | %-7s | %-7s | %s\n", customer.Name, customer.Vehicle, service.Name, transaction.Date.Format("2006-01-02"))
	}
}

func (ss *ServiceStation) GenerateCouponReport(customerID int) {
	fmt.Println("----------------------------------------")
	fmt.Println("Coupon Report - Customer ID:", customerID)
	fmt.Println("----------------------------------------")
	fmt.Println("Service | Last Coupon Date")
	fmt.Println("----------------------------------------")

	for _, service := range ss.Services {
		latestNotification := ss.getLatestNotification(customerID, service.ID)

		if !latestNotification.Date.IsZero() {
			fmt.Printf("%-7s | %s\n", service.Name, latestNotification.Date.Format("2006-01-02"))
		}
	}
}

func (ss *ServiceStation) getCustomerByID(id int) Customer {
	for _, customer := range ss.Customers {
		if customer.ID == id {
			return customer
		}
	}
	return Customer{}
}

func (ss *ServiceStation) getServiceByID(id int) Service {
	for _, service := range ss.Services {
		if service.ID == id {
			return service
		}
	}
	return Service{}
}

func (ss *ServiceStation) getLatestNotification(customerID, serviceID int) Notification {
	var latestNotification Notification

	for _, notification := range ss.Notifications {
		if notification.CustomerID == customerID && notification.ServiceID == serviceID {
			if notification.Date.After(latestNotification.Date) {
				latestNotification = notification
			}
		}
	}

	return latestNotification
}

func main() {
	// Create a new service station
	serviceStation := ServiceStation{}

	// Load data from CSV files
	err := serviceStation.LoadData()
	if err != nil {
		log.Fatal("Failed to load data:", err)
	}

	// Generate service report for a specific date
	date := time.Date(2023, 1, 1, 0, 0, 0, 0, time.UTC)
	serviceStation.GenerateServiceReport(date)

	// Generate coupon report for a specific customer
	customerID := 123
	serviceStation.GenerateCouponReport(customerID)
}
