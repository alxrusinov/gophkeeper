import styled from "styled-components";
import { FC, useContext, useState } from "react";
import { Page } from "./Page";
import { SubmitHandler, useController, useForm } from "react-hook-form";
import { User } from "src/model/user";
import TextField from "@mui/material/TextField";
import { Box, Button, Typography } from "@mui/material";
import { useMutation } from "@tanstack/react-query";
import { AxiosContext } from "../queryClient";
import { useNavigate } from "react-router-dom";
import { ErrorMessage } from "@hookform/error-message";

const Root = styled.div`
  width: 100%;
  height: 100%;
  display: grid;
  place-content: center;
`;

const Form = styled.form`
  display: flex;
  flex-direction: column;
  gap: 20px;
  align-items: center;
  justify-content: space-between;
  margin: auto;
`;

const AuthPage: FC = () => {
  const { authClient } = useContext(AxiosContext);
  const [isNewUser, setIsNewUser] = useState(false);
  const navigate = useNavigate();

  const {
    handleSubmit,
    control,
    formState: { errors },
    setError,
  } = useForm<User>({
    defaultValues: {
      username: "",
      password: "",
    },
  });

  const username = useController({
    name: "username",
    control,
    rules: { required: true },
  });

  const password = useController({
    name: "password",
    control,
    rules: { required: true },
  });

  const { mutate } = useMutation({
    mutationKey: ["auth"],
    mutationFn: async (data: User) => {
      const url = isNewUser ? "/register" : "/login";
      await authClient.post<User>(url, data);
    },
  });

  const onSubmit: SubmitHandler<User> = (data: User) => {
    mutate(data, {
      onSuccess: () => navigate("/"),
      onError: () => {
        if (isNewUser) {
          setError("username", {
            type: "server",
            message: "new user was not created",
          });
          setError("password", {
            type: "server",
            message: "new user was not created",
          });
        } else {
          setError("username", {
            type: "server",
            message: "username or password is not correct",
          });
          setError("password", {
            type: "server",
            message: "username or password is not correct",
          });
        }
      },
    });
  };

  const onToggle = () => {
    setIsNewUser((user) => !user);
  };

  console.log(errors);

  return (
    <Page>
      <Root>
        <Box
          sx={{
            width: "400px",
            paddingX: 4,
            paddingY: 6,
            borderRadius: "16px",
            boxShadow: 2,
          }}
        >
          <Typography variant="h5" gutterBottom>
            {isNewUser ? "Sign up" : "Sign in"}
          </Typography>
          <Form onSubmit={handleSubmit(onSubmit)}>
            <TextField
              value={username.field.value}
              onChange={username.field.onChange}
              name={username.field.name}
              label={"username"}
              variant="outlined"
              required
              fullWidth
            />
            <TextField
              value={password.field.value}
              onChange={password.field.onChange}
              name={password.field.name}
              label={"password"}
              variant="outlined"
              required
              fullWidth
            />
            <Box
              sx={{
                display: "flex",
                alignItems: "center",
              }}
            >
              <Typography variant="subtitle2">Or</Typography>
              <Button onClick={onToggle}>
                {isNewUser ? "Sign in" : "Sign up"}
              </Button>
            </Box>
            <ErrorMessage
              errors={errors}
              name="username"
              render={({ message }) => (
                <Typography variant="body1" color="error">
                  {message}
                </Typography>
              )}
            />
            <Button type="submit" variant="contained">
              {isNewUser ? "Sign up" : "Sign in"}
            </Button>
          </Form>
        </Box>
      </Root>
    </Page>
  );
};

export { AuthPage };
