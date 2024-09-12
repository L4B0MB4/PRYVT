import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
export const ContentCard = ({ id }: { id: string | number }) => {
  return (
    <Card key={id} className="rounded-lg shadow p-6">
      <div className="flex items-center space-x-4 mb-4">
        <Avatar>
          <AvatarImage
            src={`/placeholder.svg?height=40&width=40&text=User${id}`}
            alt={`@user${id}`}
          />
          <AvatarFallback>U{id}</AvatarFallback>
        </Avatar>
        <div>
          <h3 className="font-semibold">User {id}</h3>
          <p className="text-sm text-gray-500">2 hours ago</p>
        </div>
      </div>
      <p className="mb-4">
        This is a sample post content. It can be much longer and include various
        types of media.
      </p>
      <div className="flex justify-center">
        <img
          src="https://g-ddhl7js-ukt.vusercontent.net/placeholder.svg?height=300&width=500&text=Post+Image"
          alt="Post image"
          className="object-cover h-48 w-96 rounded-md mb-4 "
        />
      </div>
      <div className="flex space-x-4">
        <Button variant="ghost" size="sm">
          Like
        </Button>
        <Button variant="ghost" size="sm">
          Comment
        </Button>
        <Button variant="ghost" size="sm">
          Share
        </Button>
      </div>
    </Card>
  );
};
