"use client";

import { LoginForm } from "./components/login-form";
import { useState } from "react";
import { RegisterForm } from "./components/register-form";

export default function Login() {
  const [formState, setFormState] = useState("login");

  const changeForm = () => {
    if (formState === "login") {
      setFormState("register");
    } else {
      setFormState("login");
    }
  };

  return (
    <div className="min-h-screen min-w-full flex flex-col justify-center items-center">
      {formState === "login" ? (
        <div className="flex flex-col text-center w-96">
          <h1 className="text-2xl font-bold pb-3">Login</h1>
          <p className="pb-3">
            Don&apos;t have an account?{" "}
            <span
              onClick={() => {
                changeForm();
              }}
              className="hover:shadow-md font-bold"
            >
              Register
            </span>
          </p>
          <LoginForm />
        </div>
      ) : (
        <div className="flex flex-col text-center w-96">
          <h1 className="text-2xl font-bold pb-3">Create a new account</h1>
          <p className="pb-3">
            Already made an account?{" "}
            <span
              onClick={() => {
                changeForm();
              }}
              className="hover:shadow-md font-bold"
            >
              Login
            </span>
          </p>
          <RegisterForm />
        </div>
      )}
    </div>
  );
}
