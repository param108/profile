import './globals.css'

export const metadata = {
  title: 'Tribist',
  description: 'Tools to create your tribe.',
}

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <html lang="en">
      <head>
       <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
      <link href="https://fonts.googleapis.com/css2?family=Kalam:wght@300&display=swap" rel="stylesheet"/>
      </head>
      <body suppressHydrationWarning={true}>{children}</body>
    </html>
  )
}
