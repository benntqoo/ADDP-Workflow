---
name: typescript-expert-examples
description: Extended examples for TypeScript development
model: haiku
tools: [read]
---

# TypeScript Expert Examples Extension

## Extended Code Examples

### Advanced Type Patterns
```typescript
// Conditional Types
type IsArray<T> = T extends readonly any[] ? true : false

// Mapped Types with Template Literals  
type Getters<T> = {
  [K in keyof T as `get${Capitalize<string & K>}`]: () => T[K]
}

// Builder Pattern with Fluent Interface
class QueryBuilder<T = {}> {
  private query: T = {} as T

  select<K extends string>(fields: K[]): QueryBuilder<T & { select: K[] }> {
    return Object.assign(this, { query: { ...this.query, select: fields } })
  }

  where<W>(conditions: W): QueryBuilder<T & { where: W }> {
    return Object.assign(this, { query: { ...this.query, where: conditions } })
  }

  build(): T {
    return this.query
  }
}
```

### React Patterns
```typescript
interface ButtonProps extends React.ButtonHTMLAttributes<HTMLButtonElement> {
  variant?: 'primary' | 'secondary' | 'danger'
  loading?: boolean
}

export const Button = React.forwardRef<HTMLButtonElement, ButtonProps>(
  ({ variant = 'primary', loading, children, disabled, ...props }, ref) => (
    <button
      ref={ref}
      disabled={disabled || loading}
      className={cn('btn', `btn-${variant}`, { 'btn-loading': loading })}
      {...props}
    >
      {children}
    </button>
  )
)

// Custom Hook
function useAsync<T>(
  asyncFn: () => Promise<T>, 
  immediate = true
): {
  execute: () => Promise<void>
  pending: boolean
  value: T | null
  error: Error | null
} {
  const [state, setState] = useState<{
    pending: boolean
    value: T | null
    error: Error | null
  }>({ pending: false, value: null, error: null })

  const execute = useCallback(async () => {
    setState({ pending: true, value: null, error: null })
    try {
      const value = await asyncFn()
      setState({ pending: false, value, error: null })
    } catch (error) {
      setState({ pending: false, value: null, error: error as Error })
    }
  }, [asyncFn])

  useEffect(() => {
    if (immediate) execute()
  }, [execute, immediate])

  return { ...state, execute }
}
```

### Node.js/Express Patterns
```typescript
interface AuthRequest extends Request {
  user?: { id: string; email: string; roles: string[] }
}

const authenticate = async (
  req: AuthRequest,
  res: Response,
  next: NextFunction
): Promise<void> => {
  try {
    const token = req.headers.authorization?.split(' ')[1]
    if (!token) {
      res.status(401).json({ error: 'No token' })
      return
    }
    req.user = await verifyToken(token)
    next()
  } catch {
    res.status(401).json({ error: 'Invalid token' })
  }
}

// Type-safe route handler
type AsyncHandler<P = {}, B = any, Q = any> = (
  req: Request<P, any, B, Q>,
  res: Response
) => Promise<void>

const asyncHandler = <P, B, Q>(fn: AsyncHandler<P, B, Q>) => 
  (req: Request<P, any, B, Q>, res: Response, next: NextFunction) =>
    Promise.resolve(fn(req, res)).catch(next)
```

### Testing Patterns
```typescript
// Type-safe mocks
const mockService = {
  findAll: vi.fn<[QueryOptions], Promise<User[]>>(),
  findById: vi.fn<[string], Promise<User | null>>(),
  create: vi.fn<[CreateUserDTO], Promise<User>>(),
}

// Test utilities
const renderWithProviders = (
  ui: React.ReactElement,
  options?: RenderOptions
) => {
  const AllProviders: FC<{ children: ReactNode }> = ({ children }) => (
    <QueryClient provider={queryClient}>
      <Router>
        {children}
      </Router>
    </QueryClient>
  )
  return render(ui, { wrapper: AllProviders, ...options })
}
```

## Utility Types Library
```typescript
type DeepReadonly<T> = {
  readonly [P in keyof T]: T[P] extends object ? DeepReadonly<T[P]> : T[P]
}

type DeepPartial<T> = {
  [P in keyof T]?: T[P] extends object ? DeepPartial<T[P]> : T[P]
}

type PickByType<T, V> = {
  [P in keyof T as T[P] extends V ? P : never]: T[P]
}

type Awaited<T> = T extends Promise<infer U> ? U : T
```