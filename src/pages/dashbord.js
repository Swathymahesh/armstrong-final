import React, { useState } from "react";
import axios from "axios";

// Function to fetch user Armstrong numbers
export const getUserArmstrongNumbers = async (userId) => {
  const response = await axios.get(
    `http://localhost:8080/user/${userId}/numbers`
  );
  return response.data; // Return the response data
};

const UserNumbers = () => {
  const [userId, setUserId] = useState("");
  const [numbers, setNumbers] = useState([]); // State to hold fetched numbers

  // Function to fetch numbers
  const fetchNumbers = async () => {
    try {
      const result = await getUserArmstrongNumbers(userId);
      console.log("Fetched Numbers:", result); // Debug log
      setNumbers(result); // Update the numbers state
    } catch (error) {
      console.error("Error fetching numbers:", error); // Log error for debugging
      alert("Failed to fetch numbers. Please try again.");
    }
  };

  return (
    <div>
      <h1>User's Armstrong Numbers</h1>
      <input
        type="text"
        placeholder="User ID"
        value={userId}
        onChange={(e) => setUserId(e.target.value)}
      />
      <button onClick={fetchNumbers}>Fetch Numbers</button>
      <ul>
  {numbers.length > 0 ? (
    numbers.map((num) => (
      <li key={num.ID}>{num.Number || "No number available"}</li>
    ))
  ) : (
    <li>No Armstrong numbers found for this user.</li>
  )}
</ul>

    </div>
  );
};

export default UserNumbers;
