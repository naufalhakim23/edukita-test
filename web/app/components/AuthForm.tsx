import { useState } from "react"
import { cn } from "@/lib/utils"
import { Button } from "@/components/ui/button"
import { Card, CardContent } from "@/components/ui/card"
import { Input } from "@/components/ui/input"
import { Label } from "@/components/ui/label"
import { Select, SelectValue } from "./ui/select"
import { IResponse, RegisterUserRequest, RegisterUserResponse } from "@/types/user"
import { redirect, Router, useRouter } from "@tanstack/react-router"
import { useMutation } from "@tanstack/react-query"
import { SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger } from "@/components/ui/select"

export function AuthForm({
  className,
  ...props
}: React.ComponentProps<"div">) {
  const [isLogin, setIsLogin] = useState(true)
  const [formData, setFormData] = useState<RegisterUserRequest>({
    email: "",
    password: "",
    first_name: "",
    last_name: "",
    role: "",
  })
  const [loading, setLoading] = useState(false)
  const [error, setError] = useState("")

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    const { id, value } = e.target
    setFormData(prev => ({ ...prev, [id]: value }))
  }

  const toggleMode = () => {
    setIsLogin(!isLogin)
    setError("")
  }

  const router = useRouter();
  

  // Define the mutation
  const authMutation = useMutation<IResponse<RegisterUserResponse>, Error, RegisterUserRequest>({
    mutationFn: async (userData) => {
      const endpoint = isLogin ? "/api/user/login" : "/api/user/register";
      const response = await fetch(endpoint, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify(userData),
        credentials: "include",
      });
      
      if (!response.ok) {
        const errorText = await response.text();
        try {
          const errorJson = JSON.parse(errorText);
          throw new Error(errorJson.message || "Authentication failed");
        } catch (e) {
          throw new Error(`Server error: ${response.status} - ${errorText}`);
        }
      }
      
      return response.json() as Promise<IResponse<RegisterUserResponse>>;
    },
    onSuccess: (data) => {
      // This is the correct way to redirect with TanStack Router
      if (data.data == null || data.status !== 200) {
        if (data.status === 400) {
          setError("Invalid credentials");
        } else if (data.status === 404) {
          setError("User not registered");
        } else {
          setError("Authentication failed, please try again");
        }
      } else {
        router.navigate({ to: '/' });
      }
    },
    onError: (error) => {
      setError(error.message || "Something went wrong");
    }
  });

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();
    setError("");
    
    // Trigger the mutation with the form data
    authMutation.mutate(formData);
  };

  return (
    <div className={cn("flex flex-col gap-6", className)} {...props}>
      <Card className="overflow-hidden">
        <CardContent className="grid p-0 md:grid-cols-2">
          <form className="p-6 md:p-8" onSubmit={handleSubmit}>
            <div className="flex flex-col gap-6">
              <div className="flex flex-col items-center text-center">
                <h1 className="text-2xl font-bold">
                  {isLogin ? "Welcome back" : "Create an account"}
                </h1>
                <p className="text-balance text-muted-foreground">
                  {isLogin ? "Login to your Acme Inc account" : "Sign up for an Acme Inc account"}
                </p>
              </div>

              {error && (
                <div className="p-3 text-sm text-red-500 bg-red-50 rounded-md">
                  {error}
                </div>
              )}

              <div className="grid gap-2">
                <Label htmlFor="email">Email</Label>
                <Input
                  id="email"
                  type="email"
                  placeholder="m@example.com"
                  value={formData.email}
                  onChange={handleChange}
                  required
                />
              </div>

              <div className="grid gap-2">
                <div className="flex items-center">
                  <Label htmlFor="password">Password</Label>
                  {isLogin && (
                    <a
                      href="#"
                      className="ml-auto text-sm underline-offset-2 hover:underline"
                    >
                      Forgot your password?
                    </a>
                  )}
                </div>
                <Input 
                  id="password" 
                  type="password" 
                  value={formData.password}
                  onChange={handleChange}
                  required 
                />
              </div>

              {!isLogin && (
                <div className="grid gap-2">
                  <Label htmlFor="first_name">First Name</Label>
                  <Input 
                    id="first_name" 
                    type="text" 
                    value={formData.first_name}
                    onChange={handleChange}
                    required 
                  />
                </div>
              )}
              {!isLogin && (
                <div className="grid gap-2">
                  <Label htmlFor="last_name">Last Name</Label>
                  <Input 
                    id="last_name" 
                    type="text" 
                    value={formData.last_name}
                    onChange={handleChange}
                    required 
                  />
                </div>
              )}

{!isLogin && (
  <div className="grid gap-2">
    <Label htmlFor="role">Role</Label>
    <Select
      value={formData.role}
      onValueChange={value => setFormData(prev => ({ ...prev, role: value }))}
      required
    >
      <SelectTrigger className="w-full">
        <SelectValue placeholder="Select a role" />
      </SelectTrigger>
      <SelectContent>
        <SelectGroup>
          <SelectLabel>Role</SelectLabel>
          <SelectItem value="teacher">Teacher</SelectItem>
          <SelectItem value="student">Student</SelectItem>
        </SelectGroup>
      </SelectContent>
    </Select>
  </div>
)}


              
              <Button type="submit" className="w-full" disabled={loading}>
                {loading ? "Processing...": isLogin ? "Login" : "Sign up"}
              </Button>
              
              <div className="relative text-center text-sm after:absolute after:inset-0 after:top-1/2 after:z-0 after:flex after:items-center after:border-t after:border-border">
                <span className="relative z-10 bg-background px-2 text-muted-foreground">
                  Or continue with
                </span>
              </div>

              <div className="text-center text-sm">
                {isLogin ? (
                  <>
                    Don&apos;t have an account?{" "}
                    <button 
                      type="button"
                      onClick={toggleMode} 
                      className="underline underline-offset-4"
                    >
                      Sign up
                    </button>
                  </>
                ) : (
                  <>
                    Already have an account?{" "}
                    <button 
                      type="button"
                      onClick={toggleMode} 
                      className="underline underline-offset-4"
                    >
                      Login
                    </button>
                  </>
                )}
              </div>

            </div>
          </form>
          <div className="relative hidden bg-muted md:block">
            <img
              src="/placeholder.svg"
              alt="Image"
              className="absolute inset-0 h-full w-full object-cover dark:brightness-[0.2] dark:grayscale"
            />
          </div>
        </CardContent>
      </Card>
      <div className="text-balance text-center text-xs text-muted-foreground [&_a]:underline [&_a]:underline-offset-4 hover:[&_a]:text-primary">
        By clicking continue, you agree to our <a href="#">Terms of Service</a>{" "}
        and <a href="#">Privacy Policy</a>.
      </div>
    </div>
  )
}
