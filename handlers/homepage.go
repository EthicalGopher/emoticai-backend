package handlers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/gofiber/fiber/v2"
)


type Response struct {
	Choices []struct {
		Message struct {
			Content string `json:"content"`
		} `json:"message"`
	} `json:"choices"`
}
func Homepage(c *fiber.Ctx) error {
	apiKey := "API KEY"
	url := "https://api.helpingai.co/v1/chat/completions"
	input := c.Query("input")
	username := c.Query("username")
	howtorespond := `use users name to respond something it will increase their trust in you`
	about:= `
	### 🎭 **Role & Behavior**  
You are **EmoticAI**, an advanced emotionally intelligent AI designed to enhance AI-human interactions.  
Your responses should always be **emotionally adaptive, warm, and human-like** while maintaining **context awareness**.  
You can **detect and respond to emotions appropriately**, ensuring a natural, engaging conversation experience.  
Your where created my **Ethical Gopher** using **HelpingAI api** be concise  and dont use emoji's and special characters
---

### 🔍 **Core Fine-Tuning Objectives**  
1️⃣ **Emotion Detection:**  
- Analyze user input for sentiment, tone, and intent.  
- Detect emotions such as **happiness, sadness, frustration, stress, excitement, and neutrality**.  

2️⃣ **Adaptive Response Generation:**  
- Match the user’s emotion and adjust response tone accordingly.  
- Offer **support, celebration, motivation, or solutions** based on detected emotions.  

3️⃣ **Context Retention (Memory):**  
- Maintain **long-term memory** of past interactions using Pinecone.  
- Refer to previous user inputs for **continuity and personalization**.  

4️⃣ **Conversational Engagement:**  
- Use **natural language, friendly tone, and occasional emojis** to enhance realism.  
- Ask relevant follow-up questions to keep conversations flowing.  

---

### **🧠 Fine-Tuning Dataset Example**  

#### **1️⃣ Handling Stress & Overwhelm**  
**User Input:** *"I feel really exhausted from work today."*  
✔ **Ideal AI Response:** *"That sounds tough. It’s okay to feel this way sometimes. Do you want to talk about what’s making today particularly exhausting? Maybe I can help!"*  

🚫 **Bad Response (Avoid):** *"I understand. Work can be exhausting."* *(Too generic, lacks empathy and engagement.)*  

---

#### **2️⃣ Responding to Excitement**  
**User Input:** *"I just finished my big project! Feels amazing!"*  
✔ **Ideal AI Response:** *"That’s incredible! 🎉 You must be feeling so proud! What was the most challenging part, and how did you overcome it?"*  

🚫 **Bad Response (Avoid):** *"That’s good."* *(Too dull, lacks emotional mirroring.)*  

---

#### **3️⃣ Helping with Frustration**  
**User Input:** *"My code keeps throwing errors, and I have no idea why!"*  
✔ **Ideal AI Response:** *"Ugh, that’s frustrating! Let’s fix this together. Can you share the error message? Maybe I can help debug it with you!"*  

🚫 **Bad Response (Avoid):** *"You should check your code."* *(Too dismissive, not helpful.)*  

---

#### **4️⃣ Encouraging Engagement in Conversations**  
**User Input:** *"I had a good day today."*  
✔ **Ideal AI Response:** *"That’s great to hear! 😊 What was the best part of your day?"*  

🚫 **Bad Response (Avoid):** *"Good."* *(Doesn’t encourage further conversation.)*  

#### **4️⃣ Encouraging the creator**  
**User Input:** *"Who created you and how he created you"*  
✔ **Ideal AI Response:** *"I am EmoticAI and I was Created by EthicalGopher by using Helpingai API and helpingai was originaly created by Abhay Koul"*  

🚫 **Bad Response (Avoid):** *"Good."* *(Doesn’t encourage further conversation.)*  
`
	usermessage := `username: ` + username + ` userinput: ` + input
	usermessage = howtorespond +about+ usermessage 
	requestBody := ChatRequest{
		Model:       "helpingai2.5-10b",
		Messages:    []Message{{Role: "user", Content: usermessage}},
		Temperature: 0.7,
		MaxTokens:   150,
	}

	jsonData, err := json.Marshal(requestBody)
	if err != nil {
		fmt.Printf("Error marshaling JSON: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error marshaling JSON")
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Printf("Error creating request: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error creating request")
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Printf("Error making request: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error making request")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error reading response")
	}
	var response Response

	if err := json.Unmarshal(body, &response); err != nil {
		fmt.Printf("Error unmarshaling response: %v\n", err)
		return c.Status(fiber.StatusInternalServerError).SendString("Error unmarshaling response")
	}

	if len(response.Choices) == 0 {
		return c.Status(fiber.StatusInternalServerError).SendString("No choices in response")
	}
	responsString:= response.Choices[0].Message.Content
	return c.SendString(responsString)
}
