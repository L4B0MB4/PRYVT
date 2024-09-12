import { SocialNetworkLayout } from "./components/social-network-layout";
import { ThemeProvider } from "./components/theming/themeprovider";

function App() {
  return (
    <>
      <ThemeProvider defaultTheme="dark">
        <SocialNetworkLayout />
      </ThemeProvider>
    </>
  );
}

export default App;
