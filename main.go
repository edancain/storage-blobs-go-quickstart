package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/Azure/azure-sdk-for-go/sdk/azidentity"
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

const (
	account   = "https://edanstorageaccount.blob.core.windows.net/data"
	containerName = "testingcontainernames"
	blobName      = "sample-blob"
	sampleFile    = "path/to/sample/file"
)


func main() {
	// Create a default Azure credential
	cred, err := azidentity.NewDefaultAzureCredential(nil)
	if err != nil {
		log.Fatal(err)
	}

	// open the file for reading
	file, err := os.OpenFile(sampleFile, os.O_RDONLY, 0)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Create a Blob Service client
	// create a client for the specified storage account
	client, err := azblob.NewClient(account, cred, nil)
	if err != nil {
		log.Fatal(err)
	}

	ctx := context.Background()

	// upload the file to the specified container with the specified blob name
	_, err = client.UploadFile(ctx, containerName, blobName, file, nil)
	if err != nil {
		log.Fatal(err)
	}


	//DOWNLOAD A BLOB:::::::::::::::::::::::::::::::::::::
	//::::::::::::::::::::::::::::::::::::::::::::::::::::
	// this example accesses a public blob via anonymous access, so no credentials are required
	client, err = azblob.NewClientWithNoCredential("https://azurestoragesamples.blob.core.windows.net/", nil)
	if err != nil {
		log.Fatal(err)
	}

	// create or open a local file where we can download the blob
	file, err = os.Create("cloud.jpg")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// download the blob
	_, err = client.DownloadFile(context.TODO(), "samples", "cloud.jpg", file, nil)
	if err != nil {
		log.Fatal(err)
	}

	


	// ENUMERATE BLOBS

	// blob listings are returned across multiple pages
	pager := client.NewListBlobsFlatPager(containerName, nil)

	// continue fetching pages until no more remain
	for pager.More() {
	// advance to the next page
		page, err := pager.NextPage(context.TODO())
		if err != nil {
			log.Fatal(err)
		}

		// print the blob names for this page
		for _, blob := range page.Segment.BlobItems {
			fmt.Println(*blob.Name)
		}
	}
}