"'use client'";

import { Avatar, AvatarFallback, AvatarImage } from "@/components/ui/avatar";
import { Button } from "@/components/ui/button";
import { Input } from "@/components/ui/input";
import { Bell, Home, Mail, User, Users } from "lucide-react";
import { ModeToggle } from "./theming/themetoggle";

export function SocialNetworkLayout() {
  return (
    <div className="flex flex-col min-h-screen bg-gray-100">
      <header className="sticky top-0 z-10 bg-white border-b">
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
              <AvatarImage
                src="/placeholder.svg?height=32&width=32"
                alt="@user"
              />
              <AvatarFallback>U</AvatarFallback>
            </Avatar>
            <ModeToggle></ModeToggle>
          </div>
        </div>
      </header>
      <main className="flex-1">
        <div className="container mx-auto px-4 py-8">
          <div className="flex flex-col md:flex-row gap-8">
            <div className="md:w-2/3 space-y-6">
              <div className="bg-white rounded-lg shadow p-6">
                <h2 className="text-lg font-semibold mb-4">Create a post</h2>
                <textarea
                  className="w-full p-2 border border-slate-200 rounded-md dark:border-slate-800"
                  placeholder="What's on your mind?"
                  rows={3}
                />
                <div className="mt-4 flex justify-end">
                  <Button>Post</Button>
                </div>
              </div>
              {[1, 2, 3].map((post) => (
                <div key={post} className="bg-white rounded-lg shadow p-6">
                  <div className="flex items-center space-x-4 mb-4">
                    <Avatar>
                      <AvatarImage
                        src={`/placeholder.svg?height=40&width=40&text=User${post}`}
                        alt={`@user${post}`}
                      />
                      <AvatarFallback>U{post}</AvatarFallback>
                    </Avatar>
                    <div>
                      <h3 className="font-semibold">User {post}</h3>
                      <p className="text-sm text-gray-500">2 hours ago</p>
                    </div>
                  </div>
                  <p className="mb-4">
                    This is a sample post content. It can be much longer and
                    include various types of media.
                  </p>
                  <img
                    src="/placeholder.svg?height=300&width=500&text=Post+Image"
                    alt="Post image"
                    className="w-full rounded-md mb-4"
                  />
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
                </div>
              ))}
            </div>
            <div className="md:w-1/3 space-y-6">
              <div className="bg-white rounded-lg shadow p-6">
                <h2 className="text-lg font-semibold mb-4">Your Profile</h2>
                <div className="flex items-center space-x-4 mb-4">
                  <Avatar className="h-16 w-16">
                    <AvatarImage
                      src="/placeholder.svg?height=64&width=64"
                      alt="@user"
                    />
                    <AvatarFallback>U</AvatarFallback>
                  </Avatar>
                  <div>
                    <h3 className="font-semibold">John Doe</h3>
                    <p className="text-sm text-gray-500">Web Developer</p>
                  </div>
                </div>
                <Button variant="outline" className="w-full">
                  <User className="h-4 w-4 mr-2" />
                  Edit Profile
                </Button>
              </div>
              <div className="bg-white rounded-lg shadow p-6">
                <h2 className="text-lg font-semibold mb-4">
                  Suggested Connections
                </h2>
                <ul className="space-y-4">
                  {[4, 5, 6].map((user) => (
                    <li
                      key={user}
                      className="flex items-center justify-between"
                    >
                      <div className="flex items-center space-x-4">
                        <Avatar>
                          <AvatarImage
                            src={`/placeholder.svg?height=40&width=40&text=User${user}`}
                            alt={`@user${user}`}
                          />
                          <AvatarFallback>U{user}</AvatarFallback>
                        </Avatar>
                        <div>
                          <h3 className="font-semibold">User {user}</h3>
                          <p className="text-sm text-gray-500">
                            Software Engineer
                          </p>
                        </div>
                      </div>
                      <Button variant="outline" size="sm">
                        Connect
                      </Button>
                    </li>
                  ))}
                </ul>
              </div>
            </div>
          </div>
        </div>
      </main>
    </div>
  );
}
