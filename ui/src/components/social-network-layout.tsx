"'use client'";

import { Button } from "@/components/ui/button";
import { Card } from "@/components/ui/card";
import { Textarea } from "@/components/ui/textarea";
import { ContentCard } from "./contentcard/contentcard";
import { SidePanel } from "./SidePanel/SidePanel";

export function SocialNetworkLayout() {
  return (
    <div className="flex flex-col md:flex-row gap-8">
      <div className="md:w-2/3 space-y-6">
        <Card className="rounded-lg shadow p-6 ">
          <h2 className="text-lg font-semibold mb-4">Create a post</h2>
          <Textarea
            className="w-full p-2 border border-slate-200 rounded-md dark:border-slate-800"
            placeholder="What's on your mind?"
            rows={3}
          />
          <div className="mt-4 flex justify-end">
            <Button>Post</Button>
          </div>
        </Card>
        {[1, 2, 3].map((post) => (
          <ContentCard id={post} />
        ))}
      </div>
      <SidePanel />
    </div>
  );
}
