import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { User } from "lucide-react";
export const UserProfileEditPanel = () => {
  return (
    <div className="rounded-lg shadow p-6">
      <h2 className="text-lg font-semibold mb-4">Your Profile</h2>
      <div className="flex items-center space-x-4 mb-4">
        <Avatar className="h-16 w-16">
          <AvatarImage src="/placeholder.svg?height=64&width=64" alt="@user" />
          <AvatarFallback>U</AvatarFallback>
        </Avatar>
        <div>
          <h3 className="font-semibold">John Doe</h3>
          <p className="text-sm text-gray-500">Web Developer</p>
        </div>
      </div>
      <Button className="w-full">
        <User className="h-4 w-4 mr-2" />
        Edit Profile
      </Button>
    </div>
  );
};
