# YouTube Link Cleaner Bot  

A **simple Telegram bot** that removes tracking parameters from YouTube links, ensuring clean and shareable URLs.

## 🔹 How It Works  
- Detects YouTube links in messages.  
- Cleans the URL by removing unnecessary tracking parameters.  
- **Replies to the sender's message** with the cleaned YouTube link.

## 🔧 Run  
1. Get your Telegram bot token from [BotFather](https://t.me/botfather).  
2. Use TG_TOKEN enviroment variable for setting your token
   Example:
   ```sh
   TG_TOKEN="your-token" ./clean_link_bot   
   ```

