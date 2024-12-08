# cnad_assignment1

<h3>Architecture diagram design</h3>
In the diagram, when a user accesses the system through their browser, the Load Balancer manages incoming traffic to ensure high availability and reliability.
An API Gateway serves as the single entry point for all user requests. It efficiently routes these requests to the appropriate microservices based on the API endpoint, ensuring effective load distribution and enhanced security.
All microservices are hosted on Amazon EC2 instances with auto-scaling enabled. This setup dynamically creates new instances when traffic spikes or if a server goes down, maintaining availability and performance for users.
Each service is isolated, with its own dedicated microservice and corresponding database. This isolation ensures that if one service encounters an issue or fails, it does not affect the functionality of other services, promoting robust system resilience.

<h3>Architecture diagram</h3>

<h3>Instructions for setting up and running your microservices</h3>

