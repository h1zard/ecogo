// src/pages/Signup.jsx
import React, { useState } from "react";
import { useNavigate, Link } from "react-router-dom";
import { createUser } from "../api";
import { toast } from "sonner";
import "../styles.css";

const Signup = () => {
  const [form, setForm] = useState({ username: "", password: "" });
  const navigate = useNavigate();

  const handleSignup = async () => {
    try {
      await createUser(form);
      toast.success("Signup successful! You can now login.");
      navigate("/login");
    } catch (err) {
      toast.error("Signup failed");
    }
  };

  return (
    <div className="auth-container">
      <h2>Signup</h2>
      <input
        placeholder="Username"
        onChange={(e) => setForm({ ...form, username: e.target.value })}
      />
      <input
        placeholder="Password"
        type="password"
        onChange={(e) => setForm({ ...form, password: e.target.value })}
      />
      <button onClick={handleSignup}>Signup</button>
      <p>Already have an account? <Link to="/login">Login here</Link></p>
    </div>
  );
};

export default Signup;
