//code
package main

import(
	"context"
	"fmt"
	"log"
	"cloud.google.com/go/spanner"
	"google.golang.org/api/iterator"
)

const (
	projectId = "amdocs-training-489618"
	instanceId = "amdocs-tr-iiht"
	databaseId = "order-db"
)

func main(){
	ctx := context.Background()
	database := fmt.Sprintf("projects/%s/instances/%s/databases/%s",
		projectId,
		instanceId,
		databaseId,
	)
	client, err := spanner.NewClient(ctx,database)

	if err != nil {
		log.Fatalf("fail to connect spanner :%v", err)
	}

	defer client.Close()
	fmt.Println("connected to google cloud spanner....")
	/* --------------insert data-----------------
	mutations := []*spanner.Mutation{
	
		spanner.Insert(
			"Orders",
			[]string{"OrderId", "Product", "Amount", "Status", "CreatedAt"},
			[]interface{}{
				"ORD-1005",
				"Laptop",
				2000.56,
				"Pending",
				time.Now(),
			},
		),
		spanner.Insert(
			"Orders",
			[]string{"OrderId", "Product", "Amount", "Status", "CreatedAt"},
			[]interface{}{
				"ORD-1006",
				"Laptop",
				2000.56,
				"Pending",
				time.Now(),
			},
		),
		spanner.Insert(
			"Orders",
			[]string{"OrderId", "Product", "Amount", "Status", "CreatedAt"},
			[]interface{}{
				"ORD-1007",
				"Laptop",
				2000.56,
				"Pending",
				time.Now(),
			},
		),
	}
	/*
	-----------update record-------------------

	mutation := []*spanner.Mutation{
		spanner.Update(
			"Orders",
			[]string{"OrderId","Status","CreatedAt"},
			[]interface{}{
				"ORD-1006",
				"Completed",
				time.Now(),
			},
		),
	}
	

	mutation := []*spanner.Mutation{
		spanner.Delete(
			"Orders",
			spanner.KeySetFromKeys(
				spanner.Key{"ORD-1002"},
				spanner.Key{"ORD-1003"},
			),
		),
	}

	_, err = client.Apply(
		ctx,
		mutations,
	)

	*/

	rows := client.Single().Read(
		ctx,
		"Orders",
		spanner.AllKeys(),
		[]string{"OrderId", "Product", "Amount","Status"},
	)

	defer rows.Stop()

	for{
		row, err := rows.Next()

		if err == iterator.Done{
			break
		}

		if err != nil {
			log.Fatal(err)
		}

		var orderId, product, status string 
		var amount float64

		row.Columns(&orderId, &product, &amount, &status)
		if status != "Completed"{
			fmt.Println(orderId, product, amount, status)
		}
	}
}


