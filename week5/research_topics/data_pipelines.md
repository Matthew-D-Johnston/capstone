## Data Pipelines

#### What is a Data Pipeline?

A data pipeline is a means of moving data from one place (the source) to a destination (such as a data warehouse). Along the way, data is transformed and optimized, arriving in a state that can be analyzed and used to develop business insights (https://www.snowflake.com/guides/data-pipeline).

A data pipeline is any set of automated workflows that extract data from multiple sources (https://aiven.io/blog/7-challenges-that-data-pipelines-must-solve). A data pipeline essentially *is* the steps involved in aggregating, organizing, and moving data. Modern data pipelines automate many of the manual steps involved in transforming and optimizing continuous data loads (https://www.snowflake.com/guides/data-pipeline).

#### Why do Data Pipelines Exist?

Consider an organization that wants to analyze its data but those data are spread out across multiple systems and services. The data is siloed in separate data stores. Doing data analysis becomes a challenge in such a scenario (https://www.snowflake.com/guides/data-pipeline).

In order for the data to be analyzed in its entirety, the data needs to be aggregated into a single place. By consolidating data from various silos into a single source of truth, an organization is able to ensure consistent data quality and to perform quick data analysis for business insights. This is what data pipelines do (https://www.snowflake.com/guides/data-pipeline).

#### Types of Data Pipelines

There exist two main types of data pipelines, batch processing and streaming data (https://www.ibm.com/topics/data-pipeline).

##### Batch processing

As the name implies, batch processing loads “batches” of data into a repository during set time intervals, which are typically scheduled during off-peak business hours. This way, other workloads aren’t impacted as batch processing jobs tend to work with large volumes of data, which can tax the overall system. Batch processing is usually the optimal data pipeline when there isn’t an immediate need to analyze a specific dataset (e.g. monthly accounting) (https://www.ibm.com/topics/data-pipeline).  

Batch processing jobs form a workflow of sequenced commands, where the output of one command becomes the input of the next command. For example, one command may kick off data ingestion, the next command may trigger filtering of specific columns, and the subsequent command may handle aggregation. This series of commands will continue until the data is completely transformed and written into data repository (https://www.ibm.com/topics/data-pipeline).

##### Streaming data

Unlike batching processing, streaming data is leveraged when it is required for data to be continuously updated. For example, apps or point of sale systems need real-time data to update inventory and sales history of their products; that way, sellers can inform consumers if a product is in stock or not. A single action, like a product sale, is considered an “event”, and related events, such as adding an item to checkout, are typically grouped together as a “topic” or “stream.” These events are then transported via messaging systems or message brokers, such as the open-source offering, Apache Kafka (https://www.ibm.com/topics/data-pipeline).  

Since data events are processed shortly after occurring, streaming processing systems have lower latency than batch systems, but aren’t considered as reliable as batch processing systems as messages can be unintentionally dropped or spend a long time in queue. Message brokers help to address this concern through acknowledgements, where a consumer confirms processing of the message to the broker to remove it from the queue (https://www.ibm.com/topics/data-pipeline).

#### Data Pipeline Architecture

Three core steps make up the architecture of a data pipeline (https://www.ibm.com/topics/data-pipeline).

**1. Data ingestion:** Data is collected from various data sources, which includes various data structures (i.e. structured and unstructured data). Within streaming data, these raw data sources are typically known as producers, publishers, or senders. While businesses can choose to extract data only when they are ready to process it, it’s a better practice to land the raw data within a cloud data warehouse provider first. This way, the business can update any historical data if they need to make adjustments to data processing jobs (https://www.ibm.com/topics/data-pipeline). 

**2. Data Transformation:** During this step, a series of jobs are executed to process data into the format required by the destination data repository. These jobs embed automation and governance for repetitive workstreams, like business reporting, ensuring that data is cleansed and transformed consistently. For example, a data stream may come in a nested JSON format, and the data transformation stage will aim to unroll that JSON to extract the key fields for analysis (https://www.ibm.com/topics/data-pipeline).

**3. Data Storage:** The transformed data is then stored within a data repository, where it can be exposed to various stakeholders. Within streaming data, this transformed data are typically known as consumers, subscribers, or recipients (https://www.ibm.com/topics/data-pipeline). 



#### Challenges Faced by Data Pipelines

1. Getting the data where it needs to be
2. Hosting (and storing) the data
3. Flexing the data (you need to be able to flex your data today: analytics data pipes must be elastic enough to accommodate all kinds of data with few constraints on schema or type)
4. Scaling with the data
5. Moving the data around
6. Mashing up and transforming the data (means getting competitive insights on all of the data)
7. Visualizing everything (data needs to be available in a form that could be further queried and analyzed, and eventually visualized) (https://aiven.io/blog/7-challenges-that-data-pipelines-must-solve)





