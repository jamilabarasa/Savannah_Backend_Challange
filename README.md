# Savannah Backend Challenge

## functionality:

- **Google Authentication**: Login or register using Google OAuth.  
- **User Update**: Update user details like phone numbers.  
- **SMS Notification**: Send SMS notifications using Africa's Talking.  
- **Order Management**: Create orders and send SMS notifications.

---

### Base URL

The application is hosted on:  
[https://savannahbackendchallange-68506efced3c.herokuapp.com](https://savannahbackendchallange-68506efced3c.herokuapp.com)

### 1. Login/Register with Google

Visit the following endpoint to authenticate via Google:  

**URL**:  
[https://savannahbackendchallange-68506efced3c.herokuapp.com/auth/google/](https://savannahbackendchallange-68506efced3c.herokuapp.com/auth/google/)

---

### 2. Update User Phone Number

After successful authentication, update the user's phone number to enable SMS functionality.  

**Endpoint**:  
`PUT /users/:id`

**URL Example**:  
[https://savannahbackendchallange-68506efced3c.herokuapp.com/users/1](https://savannahbackendchallange-68506efced3c.herokuapp.com/users/1)

**Request Body**:  
<!-- ```json
{
  "Phone": "0741988723"
} -->


### 3. Configure SMS Notifications

Use Africa's Talking simulator to test SMS notifications.

**Africa's Talking Simulator**:  
[https://simulator.africastalking.com/](https://simulator.africastalking.com/)

**Steps**:  
1. Ensure the phone number updated in Step 2 is correctly set for the user.  
2. Use the Africa's Talking simulator to verify SMS delivery before sending sms notifications.  

---

### 4. Create an Order

Place an order and trigger an SMS notification to the customer.

**Endpoint**:  
`POST /orders`

**URL**:  
[https://savannahbackendchallange-68506efced3c.herokuapp.com/orders](https://savannahbackendchallange-68506efced3c.herokuapp.com/orders)

**Request Body**:  
```json
{
  "item": "test",
  "amount": 50.00,
  "time": "2024-11-19T13:17:28.004831698+03:00",
  "customerid": 1
}


## Notes
- 
- All requests must include proper authentication via the Google OAuth endpoint described in Step 1.  
-Use Customer ID from auth/login callback response
- The `customerid` in the request body must correspond to an existing user in the database.  
- Use the Africa's Talking simulator to test SMS functionality before deploying the system live.  
- Ensure all request bodies are correctly formatted as JSON to avoid errors.  

---

# How to Run Locally

### Prerequisites

- [Go 1.18 or higher](https://go.dev/dl/) 
- Africa's Talking SMS account config  
- A Google console Cloud config

### Steps

1. **Clone the Repository**  
   ```bash
   git clone https://github.com/jamilabarasa/Savannah_Backend_Challange.git



2. **Set Up Environment Variables**
GOOGLE_CLIENT_ID=
GOOGLE_CLIENT_SECRET=
JWT_SECRET=
DATABASE_URL=
AFRICA_TALKING_API_KEY=
AFRICA_TALKING_USERNAME=
AFRICA_TALKING_API_URL=
AFRICA_TALKING_SHORTCODE=
PORT=3000

3. **Install Dependencies**
   ```bash
   go mod tidy
   
4. **Start the Application**
   ```bash
   go run main.go
