package alienvault

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestAWSBucketJob(t *testing.T) {

	testJob := AWSBucketJob{
		Params: AWSBucketJobParams{
			BucketName: "my-bucket",
			Path:       "/logs",
		},
	}

	// promoted fields
	testJob.Name = "test-client-my-bucket-job"
	testJob.Description = "This is an auto-generated test job made by https://github.com/form3tech-oss/alienvault"
	testJob.SensorID = "aaaaaaaa-aaaa-aaaa-aaaaaaaaaaaa"
	testJob.Schedule = JobScheduleHourly

	// promoted params fields
	testJob.Params.Plugin = "PostgreSQL"
	testJob.Params.SourceFormat = JobSourceFormatRaw

	// test creating

	if err := testClient.CreateAWSBucketJob(&testJob); err != nil {
		t.Fatalf("Failed to create job: %s", err)
	}

	require.NotEmpty(t, testJob.UUID, "A created job should be assigned a UUID")

	// test reading

	refreshedJob, err := testClient.GetAWSBucketJob(testJob.UUID)
	if err != nil {
		t.Fatalf("Failed to refresh job: %s", err)
	}

	assert.Equal(t, refreshedJob.Name, testJob.Name, "Job fields should be set")
	assert.Equal(t, refreshedJob.Description, testJob.Description, "Job fields should be set")
	assert.Equal(t, refreshedJob.SensorID, testJob.SensorID, "Job fields should be set")
	assert.Equal(t, refreshedJob.Params.SourceFormat, testJob.Params.SourceFormat, "Job fields should be set")
	assert.Equal(t, refreshedJob.Params.Plugin, testJob.Params.Plugin, "Job fields should be set")
	assert.Equal(t, refreshedJob.Params.BucketName, testJob.Params.BucketName, "Job fields should be set")
	assert.Equal(t, refreshedJob.Params.Path, testJob.Params.Path, "Job fields should be set")

	// test list jobs
	jobs, err := testClient.GetAWSBucketJobs()
	if err != nil {
		t.Fatalf("Failed to list jobs: %s", err)
	}

	found := false
	for _, job := range jobs {
		if job.UUID == testJob.UUID {
			found = true
			break
		}
	}
	assert.True(t, found, "Created jobs must show up in the job list")

	// test updating

	testJob.Name = testJob.Name + "-updated"
	testJob.Params.Plugin = "Nginx"
	testJob.Params.BucketName = "updated-bucket-name"

	if err := testClient.UpdateAWSBucketJob(&testJob); err != nil {
		t.Fatalf("Failed to update job: %s", err)
	}

	refreshedJob, err = testClient.GetAWSBucketJob(testJob.UUID)
	if err != nil {
		t.Fatalf("Failed to refresh job: %s", err)
	}

	assert.Equal(t, refreshedJob.Name, testJob.Name, "Job fields should be updated")
	assert.Equal(t, refreshedJob.Description, testJob.Description, "Job fields should be updated")
	assert.Equal(t, refreshedJob.SensorID, testJob.SensorID, "Job fields should be updated")
	assert.Equal(t, refreshedJob.Params.SourceFormat, testJob.Params.SourceFormat, "Job fields should be updated")
	assert.Equal(t, refreshedJob.Params.Plugin, testJob.Params.Plugin, "Job fields should be updated")
	assert.Equal(t, refreshedJob.Params.BucketName, testJob.Params.BucketName, "Job fields should be updated")
	assert.Equal(t, refreshedJob.Params.Path, testJob.Params.Path, "Job fields should be updated")

	// test deleting

	if err := testClient.DeleteAWSBucketJob(&testJob); err != nil {
		t.Fatalf("Failed to delete job: %s", err)
	}

	if _, err := testClient.GetAWSBucketJob(testJob.UUID); err == nil {
		t.Fatalf("Job still exists after deletion")
	}

}
