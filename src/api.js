import axios from "axios";

// Set the base URL for your backend
const API_URL = "http://localhost:8080";

// Create an Axios instance for consistent configurations
const axiosInstance = axios.create({
  baseURL: API_URL,
  headers: {
    "Content-Type": "application/json", // Ensure the correct header is always sent
  },
});

// Function to register a user
export const registerUser = async (email) => {
  try {
    const response = await axiosInstance.post("/register", { email });
    return response.data;
  } catch (error) {
    console.error("Error during user registration:", error);
    throw error.response ? error.response.data : error.message;
  }
};

// Function to verify an Armstrong number
export const verifyArmstrong = async (userId, number) => {
  try {
    const response = await axiosInstance.post("/verify", {
      user_id: userId,
      number: Number(number)
    });
    return response.data;
  } catch (error) {
    console.error("Error during Armstrong number verification:", error);
    throw error.response ? error.response.data : error.message;
  }
};

// Function to fetch a user's Armstrong numbers
export const getUserArmstrongNumbers = async (userId) => {
  try {
    const response = await axiosInstance.get(`/user/${userId}/numbers`);
    return response.data;
  } catch (error) {
    console.error("Error fetching user's Armstrong numbers:", error);
    throw error.response ? error.response.data : error.message;
  }
};

// Function to fetch all users and their Armstrong numbers
export const getAllUsers = async () => {
  try {
    const response = await axiosInstance.get("/users");
    return response.data;
  } catch (error) {
    console.error("Error fetching all users and their Armstrong numbers:", error);
    throw error.response ? error.response.data : error.message;
  }
};
