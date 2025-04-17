import { SignUpForm } from '@/components/SignUpForm'
import { createFileRoute, useRouter } from '@tanstack/react-router'

export const Route = createFileRoute('/signup')({
  component: SignupPage,
})

function SignupPage() {
  return (<div className="flex min-h-svh flex-col items-center justify-center bg-muted p-6 md:p-10">
        <div className="w-full max-w-sm md:max-w-3xl">
          <SignUpForm />
        </div>
      </div>)
}
