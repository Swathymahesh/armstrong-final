import React, { useState } from "react";
import "./home.css";
import axios from "axios";


const Home = () => {
  const [formValues, setFormValues] = useState({
    name: "",
    email: "",
    password: "",
  });

  const [formErrors, setFormErrors] = useState({});
  const [isSubmitted, setIsSubmitted] = useState(false);
  const [apiResponse, setApiResponse] = useState("");

  const handleChange = (e) => {
    const { name, value } = e.target;
    setFormValues({ ...formValues, [name]: value });
  };

  const validate = (values) => {
    const errors = {};
    const emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;

    if (!values.name.trim()) {
      errors.name = "Name is required";
    }

    if (!values.email) {
      errors.email = "Email is required";
    } else if (!emailRegex.test(values.email)) {
      errors.email = "Email is invalid";
    }

    if (!values.password) {
      errors.password = "Password is required";
    } else if (values.password.length < 6) {
      errors.password = "Password must be at least 6 characters long";
    }

    return errors;
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    const errors = validate(formValues);
    setFormErrors(errors);

    if (Object.keys(errors).length === 0) {
      try {
        // Call the backend API
        const response = await axios.post("http://localhost:8080/register", {
          email: formValues.email,
         // password: formValues.password,
        });

        setIsSubmitted(true);
        setApiResponse("Form submitted successfully: " + response.data);
      } catch (error) {
        setApiResponse(
          error.response
            ? `Error: ${error.response.data}`
            : "An error occurred. Please try again."
        );
        setIsSubmitted(false);
      }
    } else {
      setIsSubmitted(false);
    }
  };

  return (
    <div className="home">
      <header className="home-header">
        <h1>Validation Form</h1>
        <p>Please fill in the details below:</p>
      </header>

      {isSubmitted ? (
        <div className="success-message">
          {apiResponse || "Form submitted successfully!"}
        </div>
      ) : (
        <form className="form" onSubmit={handleSubmit}>
          <div className="form-group">
            <label htmlFor="name">Name:</label>
            <input
              type="text"
              id="name"
              name="name"
              value={formValues.name}
              onChange={handleChange}
              placeholder="Enter your name"
            />
            {formErrors.name && <p className="error">{formErrors.name}</p>}
          </div>

          <div className="form-group">
            <label htmlFor="email">Email:</label>
            <input
              type="email"
              id="email"
              name="email"
              value={formValues.email}
              onChange={handleChange}
              placeholder="Enter your email"
            />
            {formErrors.email && <p className="error">{formErrors.email}</p>}
          </div>

          <div className="form-group">
            <label htmlFor="password">Password:</label>
            <input
              type="password"
              id="password"
              name="password"
              value={formValues.password}
              onChange={handleChange}
              placeholder="Enter your password"
            />
            {formErrors.password && <p className="error">{formErrors.password}</p>}
          </div>


          <button type="submit" className="submit-btn">
            Submit
          </button>
        </form>
      )}

      {apiResponse && <div className="api-response">{apiResponse}</div>}
    </div>
  );
};

export default Home;
