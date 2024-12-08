# CNAD Assignment 1

## Architecture Diagram Design

In the architecture design:

1. **Load Balancer**  
   - Manages incoming traffic to ensure high availability and reliability.

2. **API Gateway**  
   - Acts as the single entry point for all user requests.  
   - Routes requests to the appropriate microservices based on the API endpoint, ensuring effective load distribution and enhanced security.

3. **Microservices on Amazon EC2**  
   - All microservices are hosted on Amazon EC2 instances with auto-scaling enabled.  
   - Auto-scaling ensures new instances are created dynamically during traffic spikes or server failures, maintaining system availability and performance.

4. **Service Isolation**  
   - Each service operates independently with its own dedicated microservice and database.  
   - This design ensures that a failure in one service does not impact the functionality of others, promoting robust system resilience.

---

## Instructions for Setting Up and Running Microservices

Follow these steps to set up and run your microservices:

1. **Database Setup**  
   - Import the database files (`billing_db.sql`, `reservation_db.sql`, `user_db.sql`, `vehicle_db.sql`) into **MySQL Workbench**.

2. **Update Database Credentials**  
   - Modify the database connection credentials in the following files:
     - **`billing_db.go`** (located in `assignment1/services/billing/utils`)
     - **`reservation_db.go`** (located in `assignment1/services/reservation/utils`)
     - **`db.go`** (located in `assignment1/services/user/utils`)
     - **`vehicle_db.go`** (located in `assignment1/services/vehicle/utils`)
   - Example modification:
     ```go
     db, err = sql.Open("mysql", "root:04D685362v98@tcp(127.0.0.1:3306)/billing_db")
     ```

3. **Update Main Service Endpoint**  
   - In `main.go`, change the URL for the localhost (`localhost:8080`) to your specific localhost URL.

---

