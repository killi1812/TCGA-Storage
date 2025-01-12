package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Insert(patient PatientData) error {
	dbLock.Lock()
	defer dbLock.Unlock()

	result, err := patients.InsertOne(context.Background(), patient)
	if err != nil {
		return err
	}

	fmt.Printf("result: %v\n", result)

	return nil
}

func InsertMany(arr []PatientData) error {
	dbLock.Lock()
	defer dbLock.Unlock()

	bson := make([]interface{}, len(arr))
	for i := 0; i < len(arr); i++ {
		bson[i] = arr[i]
	}

	_, err := patients.InsertMany(context.Background(), bson)
	if err != nil {
		return err
	}

	return nil
}

func Read(patientCode string) (PatientData, error) {
	//TODO:
	//check if lock needed
	//need to wait utill unlocked
	//dbLock.Lock()
	//defer dbLock.Unlock()

	var patient PatientData
	filter := bson.M{"bcr_patient_barcode": patientCode}

	err := patients.FindOne(context.Background(), filter).Decode(&patient)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			fmt.Printf("patient: %v\n", patient)
			return PatientData{}, fmt.Errorf("no patient found with code %s", patientCode)
		}
		return PatientData{}, err
	}

	return patient, nil
}

func ReadAll() ([]PatientData, error) {
	cursor, err := patients.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}

	var result []PatientData
	if err = cursor.All(context.Background(), &result); err != nil {
		return nil, err
	}

	return result, nil
}

func DeleteAll() error {
	_, err := patients.DeleteMany(context.Background(), bson.D{})
	if err != nil {
		return err
	}

	return nil
}
