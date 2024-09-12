"'use client'";

import { SuggestedUser } from "../suggesteduser/SuggestedUser";
import { UserProfileEditPanel } from "../UserProfileEdit/UserProfileEdit";

export const SidePanel = () => {
  return (
    <div className="md:w-1/3 space-y-6">
      <UserProfileEditPanel />
      <div className="rounded-lg shadow p-6">
        <h2 className="text-lg font-semibold mb-4">Suggested Connections</h2>
        <ul className="space-y-4">
          {[4, 5, 6].map((user) => (
            <SuggestedUser userId={user + ""} />
          ))}
        </ul>
      </div>
    </div>
  );
};
