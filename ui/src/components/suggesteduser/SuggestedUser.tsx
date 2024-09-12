import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
export const SuggestedUser = ({ userId }: { userId: string }) => {
  return (
    <li key={userId} className="flex items-center justify-between">
      <div className="flex items-center space-x-4">
        <Avatar>
          <AvatarImage
            src={`/placeholder.svg?height=40&width=40&text=User${userId}`}
            alt={`@user${userId}`}
          />
          <AvatarFallback>U{userId}</AvatarFallback>
        </Avatar>
        <div>
          <h3 className="font-semibold">User {userId}</h3>
          <p className="text-sm text-gray-500">Software Engineer</p>
        </div>
      </div>
      <Button size="sm">Connect</Button>
    </li>
  );
};
