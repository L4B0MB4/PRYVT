import { Header } from "./components/header/header";
import { SocialNetworkLayout } from "./components/social-network-layout";
import { ThemeProvider } from "./components/theming/themeprovider";

function App() {
  return (
    <>
      <ThemeProvider defaultTheme="dark">
        <div className="bg-background">
          <div className="flex flex-col min-h-screen">
            <Header />
            <main className="flex-1">
              <div className="container mx-auto px-4 py-8">
                <SocialNetworkLayout />
              </div>
            </main>
          </div>
        </div>
      </ThemeProvider>
    </>
  );
}

export default App;
