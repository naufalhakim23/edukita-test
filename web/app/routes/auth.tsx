import { AuthForm } from "@/components/AuthForm"
import { createFileRoute, redirect } from '@tanstack/react-router'


export const Route = createFileRoute('/auth')({
  component: LoginPage,
})

function LoginPage() {
  return (
    <div className="flex min-h-svh flex-col items-center justify-center bg-muted p-6 md:p-10">
      <div className="w-full max-w-sm md:max-w-3xl">
        <AuthForm />
      </div>
    </div>
  )
}
