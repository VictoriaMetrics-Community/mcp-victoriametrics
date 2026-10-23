import * as React from "react"
import { useState, useCallback } from "react"
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from "@/components/ui/card"
import { Button } from "@/components/ui/button"
import { Badge } from "@/components/ui/badge"
import { Input } from "@/components/ui/input"
import { Collapsible, CollapsibleContent, CollapsibleTrigger } from "@/components/ui/collapsible"
import { Copy, Check, ChevronDown, Wind } from "lucide-react"
import {ClaudeAI, Cursor, Gemini, JetBrains, OpenAI, OpenCode, VisualStudioCode, Zed} from "@/components/tabs/icons.tsx";
import clients from "@/clients.gen.json"

const clientIcons: Record<string, React.ComponentType<{ className?: string }>> = {
  "claude-desktop": ClaudeAI,
  "claude-code": ClaudeAI,
  "cursor": Cursor,
  "codex": OpenAI,
  "vscode": VisualStudioCode,
  "zed": Zed,
  "jetbrains": JetBrains,
  "windsurf": Wind,
  "gemini-cli": Gemini,
  "opencode": OpenCode,
}

interface ConfigItemProps {
  title: string
  description: string
  icon: React.ReactNode
  config: string
  badge?: string
  defaultOpen?: boolean
}

function ConfigItem({ title, description, icon, config, badge, defaultOpen = false }: ConfigItemProps) {
  const [copied, setCopied] = useState(false)
  const [isOpen, setIsOpen] = useState(defaultOpen)

  const handleCopy = useCallback(async () => {
    await navigator.clipboard.writeText(config)
    setCopied(true)
    setTimeout(() => setCopied(false), 2000)
  }, [config])

  return (
    <Collapsible open={isOpen} onOpenChange={setIsOpen}>
      <Card>
        <CollapsibleTrigger asChild>
          <CardHeader className="pb-3 cursor-pointer hover:bg-muted/50 transition-colors">
            <div className="flex items-center justify-between">
              <div className="flex items-center gap-3">
                <div className="flex h-10 w-10 items-center justify-center rounded-lg bg-muted">
                  {icon}
                </div>
                <div>
                  <CardTitle className="text-base flex items-center gap-2">
                    {title}
                    {badge && <Badge variant="secondary">{badge}</Badge>}
                  </CardTitle>
                  <CardDescription className="text-sm">{description}</CardDescription>
                </div>
              </div>
              <ChevronDown className={`h-5 w-5 text-muted-foreground transition-transform duration-200 ${isOpen ? 'rotate-180' : ''}`} />
            </div>
          </CardHeader>
        </CollapsibleTrigger>
        <CollapsibleContent>
          <CardContent>
            <div className="relative">
              <pre className="overflow-x-auto rounded-lg bg-slate-900 p-4 text-sm text-slate-100">
                <code>{config}</code>
              </pre>
              <Button
                variant="secondary"
                size="sm"
                className="absolute right-2 top-2"
                onClick={handleCopy}
              >
                {copied ? (
                  <>
                    <Check className="h-4 w-4" />
                    Copied
                  </>
                ) : (
                  <>
                    <Copy className="h-4 w-4" />
                    Copy
                  </>
                )}
              </Button>
            </div>
          </CardContent>
        </CollapsibleContent>
      </Card>
    </Collapsible>
  )
}

export function ClientConfigs() {
  const [serverUrl, setServerUrl] = useState(() => {
    if (typeof window !== 'undefined') {
      return window.location.origin
    }
    return 'http://localhost:8080'
  })

  const configs: ConfigItemProps[] = clients.map((client, i) => {
    const Icon = clientIcons[client.id] ?? Wind
    return {
      title: client.title,
      description: client.description,
      icon: <Icon className="h-6 w-6" />,
      badge: "badge" in client ? client.badge : undefined,
      defaultOpen: i === 0,
      config: client.remote.code.replaceAll("{{url}}", serverUrl).trimEnd(),
    }
  })

  return (
    <div className="space-y-6">
      <div className="space-y-2">
        <h2 className="text-2xl font-bold tracking-tight">Client Configuration</h2>
        <p className="text-muted-foreground">
          Copy these configuration snippets to connect your favorite MCP client to this server.
        </p>
      </div>

      <Card>
        <CardHeader className="pb-3">
          <CardTitle className="text-base">Server URL</CardTitle>
          <CardDescription>
            The base URL of this MCP server. Update if accessing from a different host.
          </CardDescription>
        </CardHeader>
        <CardContent>
          <div className="flex gap-2">
            <Input
              value={serverUrl}
              onChange={(e) => setServerUrl(e.target.value)}
              placeholder="http://localhost:8080"
              className="font-mono"
            />
            <Button
              variant="secondary"
              onClick={() => setServerUrl(window.location.origin)}
            >
              Reset
            </Button>
          </div>
        </CardContent>
      </Card>

      <div className="grid gap-4">
        {configs.map((config) => (
          <ConfigItem key={config.title} {...config} />
        ))}
      </div>
    </div>
  )
}
