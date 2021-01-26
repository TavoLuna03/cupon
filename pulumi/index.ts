import * as aws from "@pulumi/aws";

module.exports = function () {
    const dynamodb_items_test = new aws.dynamodb.Table(`items-test`, {
        name: `items-test`,
        billingMode: "PAY_PER_REQUEST",
        attributes: [
            { name: "id", type: "S" },
        ],
        hashKey: "id",
    });

    return {
        dynamo_db_items_testarn: dynamodb_items_test.arn,
    }
}