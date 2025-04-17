import { createFileRoute } from '@tanstack/react-router'

export const Route = createFileRoute('/assignments')({
  component: RouteComponent,
})

function RouteComponent() {
  return <div>Hello "/assignments"!</div>
}
