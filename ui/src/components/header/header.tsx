import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Bell, Home, Mail, Users } from "lucide-react";
import { ModeToggle } from "../theming/themetoggle";
export const Header = () => {
  return (
    <header className="sticky top-0 z-10 border-b bg-background">
      <div className="container mx-auto px-4 py-3 flex items-center justify-between">
        <div className="flex items-center space-x-4">
          <h1 className="text-2xl font-bold text-slate-900 dark:text-slate-50">
            SocialNet
          </h1>
          <nav className="hidden md:flex space-x-4">
            <Button variant="ghost" size="sm">
              <Home className="h-4 w-4 mr-2" />
              Home
            </Button>
            <Button variant="ghost" size="sm">
              <Users className="h-4 w-4 mr-2" />
              Network
            </Button>
            <Button variant="ghost" size="sm">
              <Mail className="h-4 w-4 mr-2" />
              Messages
            </Button>
          </nav>
        </div>
        <div className="flex items-center space-x-4">
          <form className="hidden md:block">
            <Input type="search" placeholder="Search..." className="w-64" />
          </form>
          <Button variant="ghost" size="icon">
            <Bell className="h-5 w-5" />
          </Button>
          <Avatar>
            <AvatarImage alt="@user" />
            <AvatarFallback>U</AvatarFallback>
          </Avatar>
          <ModeToggle></ModeToggle>
        </div>
      </div>
    </header>
  );
};
